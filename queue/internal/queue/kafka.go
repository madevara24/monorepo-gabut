package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"

	"queue/internal/config"
)

// kafkaQueue uses one topic and one consumer group.
// Kafka stores messages in the topic after the consumer reads them.
// The consumer commits an offset to mark a message as done.
// A handler error leaves the offset uncommitted, so the message comes back.
type kafkaQueue struct {
	brokers []string
	topic   string
	group   string
	writer  *kafka.Writer
}

func newKafka(cfg config.Config) (*kafkaQueue, error) {
	if err := ensureTopic(cfg.Kafka.Brokers, cfg.Kafka.Topic); err != nil {
		return nil, err
	}

	writer := &kafka.Writer{
		Addr:                   kafka.TCP(cfg.Kafka.Brokers...),
		Topic:                  cfg.Kafka.Topic,
		Balancer:               &kafka.LeastBytes{},
		RequiredAcks:           kafka.RequireOne,
		AllowAutoTopicCreation: true,
	}

	return &kafkaQueue{
		brokers: cfg.Kafka.Brokers,
		topic:   cfg.Kafka.Topic,
		group:   cfg.Kafka.Group,
		writer:  writer,
	}, nil
}

func (q *kafkaQueue) Publish(ctx context.Context, msg Message) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return q.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(msg.ID),
		Value: body,
	})
}

func (q *kafkaQueue) Subscribe(ctx context.Context, handler Handler) error {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  q.brokers,
		Topic:    q.topic,
		GroupID:  q.group,
		MinBytes: 1,
		MaxBytes: 10 << 20,
	})
	defer reader.Close()

	for {
		m, err := reader.FetchMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return nil
			}
			return fmt.Errorf("kafka fetch: %w", err)
		}

		msg, err := decodeMessage(m.Value)
		if err != nil {
			fmt.Fprintf(os.Stderr, "kafka: skip bad message at offset %d: %v\n", m.Offset, err)
			reader.CommitMessages(ctx, m)
			continue
		}

		if err := handler(ctx, msg); err != nil {
			fmt.Fprintf(os.Stderr, "kafka: handler error, offset %d stays uncommitted: %v\n", m.Offset, err)
			continue
		}

		if err := reader.CommitMessages(ctx, m); err != nil {
			return fmt.Errorf("kafka commit: %w", err)
		}
	}
}

func (q *kafkaQueue) Close() error {
	return q.writer.Close()
}

// ensureTopic creates the topic when it is absent.
// A consumer group needs the topic before it can join the group.
func ensureTopic(brokers []string, topic string) error {
	if len(brokers) == 0 {
		return fmt.Errorf("kafka: no brokers configured")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := kafka.DialContext(ctx, "tcp", brokers[0])
	if err != nil {
		return fmt.Errorf("kafka dial %s: %w", brokers[0], err)
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return fmt.Errorf("kafka controller lookup: %w", err)
	}

	addr := net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port))
	ctrl, err := kafka.DialContext(ctx, "tcp", addr)
	if err != nil {
		return fmt.Errorf("kafka controller dial %s: %w", addr, err)
	}
	defer ctrl.Close()

	// CreateTopics reports an error when the topic already exists. That error is safe to ignore.
	_ = ctrl.CreateTopics(kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     1,
		ReplicationFactor: 1,
	})

	if _, err := ctrl.ReadPartitions(topic); err != nil {
		return fmt.Errorf("kafka topic %q is not ready: %w", topic, err)
	}
	return nil
}
