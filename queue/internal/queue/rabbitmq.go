package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"

	"queue/internal/config"
)

// rabbitQueue uses one durable queue.
// The consumer acknowledges every message.
// A handler error sends the message back to the queue for redelivery.
type rabbitQueue struct {
	conn  *amqp.Connection
	pub   *amqp.Channel
	sub   *amqp.Channel
	queue string
}

func newRabbitMQ(cfg config.Config) (*rabbitQueue, error) {
	conn, err := amqp.Dial(cfg.RabbitMQ.URL)
	if err != nil {
		return nil, fmt.Errorf("rabbitmq dial: %w", err)
	}

	pub, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("rabbitmq publish channel: %w", err)
	}

	sub, err := conn.Channel()
	if err != nil {
		pub.Close()
		conn.Close()
		return nil, fmt.Errorf("rabbitmq consume channel: %w", err)
	}

	// durable=true keeps the queue after a broker restart.
	if _, err := sub.QueueDeclare(cfg.RabbitMQ.Queue, true, false, false, false, nil); err != nil {
		sub.Close()
		pub.Close()
		conn.Close()
		return nil, fmt.Errorf("rabbitmq declare queue: %w", err)
	}

	return &rabbitQueue{conn: conn, pub: pub, sub: sub, queue: cfg.RabbitMQ.Queue}, nil
}

func (q *rabbitQueue) Publish(ctx context.Context, msg Message) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return q.pub.PublishWithContext(ctx, "", q.queue, false, false, amqp.Publishing{
		ContentType:  "application/json",
		DeliveryMode: amqp.Persistent,
		Body:         body,
	})
}

func (q *rabbitQueue) Subscribe(ctx context.Context, handler Handler) error {
	deliveries, err := q.sub.Consume(q.queue, "", false, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("rabbitmq consume: %w", err)
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case d, ok := <-deliveries:
			if !ok {
				return nil
			}
			msg, err := decodeMessage(d.Body)
			if err != nil {
				fmt.Fprintf(os.Stderr, "rabbitmq: skip bad message: %v\n", err)
				d.Nack(false, false)
				continue
			}
			if err := handler(ctx, msg); err != nil {
				fmt.Fprintf(os.Stderr, "rabbitmq: handler error, requeue: %v\n", err)
				d.Nack(false, true) // requeue=true sends the message back
				continue
			}
			d.Ack(false)
		}
	}
}

func (q *rabbitQueue) Close() error {
	q.pub.Close()
	q.sub.Close()
	return q.conn.Close()
}
