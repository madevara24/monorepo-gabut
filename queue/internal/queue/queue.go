package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"queue/internal/config"
)

// Message is one unit that travels through a queue.
// Body is raw JSON. The sender copies it from the payload template.
type Message struct {
	ID     string          `json:"id"`
	SentAt time.Time       `json:"sentAt"`
	Body   json.RawMessage `json:"body"`
}

// Handler processes one message.
// A nil return tells the queue to acknowledge the message.
// A non-nil return tells the queue to keep or requeue the message for redelivery.
type Handler func(ctx context.Context, msg Message) error

// Queue is the one interface that every broker implements.
// The send, receive, and demo commands use this interface only.
type Queue interface {
	Publish(ctx context.Context, msg Message) error
	Subscribe(ctx context.Context, handler Handler) error
	Close() error
}

// New returns the queue implementation for the configured broker.
func New(cfg config.Config) (Queue, error) {
	switch cfg.Broker {
	case "redis":
		return newRedis(cfg)
	case "rabbitmq":
		return newRabbitMQ(cfg)
	case "kafka":
		return newKafka(cfg)
	default:
		return nil, fmt.Errorf("unknown broker %q", cfg.Broker)
	}
}

func decodeMessage(data []byte) (Message, error) {
	var msg Message
	if err := json.Unmarshal(data, &msg); err != nil {
		return Message{}, err
	}
	return msg, nil
}
