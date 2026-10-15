package queue

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"

	"queue/internal/config"
)

// redisQueue has two modes.
//
// Streams mode uses a consumer group. The consumer acknowledges every message.
// An unacknowledged message stays in the stream and comes back later.
//
// List mode uses LPUSH and BRPOP. A pop removes the message at once.
// A consumer that crashes after the pop loses that message.
type redisQueue struct {
	client   *redis.Client
	mode     string
	stream   string
	group    string
	consumer string
}

func newRedis(cfg config.Config) (*redisQueue, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis ping %s: %w", cfg.Redis.Addr, err)
	}

	host, _ := os.Hostname()
	return &redisQueue{
		client:   client,
		mode:     cfg.Redis.Mode,
		stream:   cfg.Redis.Stream,
		group:    cfg.Redis.Group,
		consumer: fmt.Sprintf("%s-%d", host, os.Getpid()),
	}, nil
}

func (q *redisQueue) Publish(ctx context.Context, msg Message) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	if q.mode == "list" {
		return q.client.LPush(ctx, q.stream, string(data)).Err()
	}
	return q.client.XAdd(ctx, &redis.XAddArgs{
		Stream: q.stream,
		Values: map[string]any{"data": string(data)},
	}).Err()
}

func (q *redisQueue) Subscribe(ctx context.Context, handler Handler) error {
	if q.mode == "list" {
		return q.subscribeList(ctx, handler)
	}
	return q.subscribeStreams(ctx, handler)
}

func (q *redisQueue) subscribeList(ctx context.Context, handler Handler) error {
	for {
		if ctx.Err() != nil {
			return nil
		}
		result, err := q.client.BRPop(ctx, 0, q.stream).Result()
		if err != nil {
			if errors.Is(err, redis.Nil) {
				continue
			}
			if ctx.Err() != nil {
				return nil
			}
			return err
		}
		msg, err := decodeMessage([]byte(result[1]))
		if err != nil {
			fmt.Fprintf(os.Stderr, "redis list: skip bad message: %v\n", err)
			continue
		}
		if err := handler(ctx, msg); err != nil {
			fmt.Fprintf(os.Stderr, "redis list: handler error: %v\n", err)
		}
	}
}

func (q *redisQueue) subscribeStreams(ctx context.Context, handler Handler) error {
	// "$" means the group reads only messages that arrive after this call.
	err := q.client.XGroupCreateMkStream(ctx, q.stream, q.group, "$").Err()
	if err != nil && !strings.Contains(err.Error(), "BUSYGROUP") {
		return fmt.Errorf("redis create group: %w", err)
	}

	for {
		if ctx.Err() != nil {
			return nil
		}
		streams, err := q.client.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    q.group,
			Consumer: q.consumer,
			Streams:  []string{q.stream, ">"},
			Count:    10,
			Block:    5 * time.Second,
		}).Result()
		if err != nil {
			if errors.Is(err, redis.Nil) {
				continue
			}
			if ctx.Err() != nil {
				return nil
			}
			return err
		}

		for _, stream := range streams {
			for _, entry := range stream.Messages {
				raw, _ := entry.Values["data"].(string)
				msg, err := decodeMessage([]byte(raw))
				if err != nil {
					fmt.Fprintf(os.Stderr, "redis stream: skip bad entry %s: %v\n", entry.ID, err)
					q.client.XAck(ctx, q.stream, q.group, entry.ID)
					continue
				}
				if err := handler(ctx, msg); err != nil {
					fmt.Fprintf(os.Stderr, "redis stream: handler error, entry %s stays unacked: %v\n", entry.ID, err)
					continue
				}
				q.client.XAck(ctx, q.stream, q.group, entry.ID)
			}
		}
	}
}

func (q *redisQueue) Close() error {
	return q.client.Close()
}
