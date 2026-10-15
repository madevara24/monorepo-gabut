package cmd

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/spf13/cobra"

	"queue/internal/payload"
	"queue/internal/queue"
)

var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "Publish and consume in one process",
	RunE: func(_ *cobra.Command, _ []string) error {
		ctx, stop := commandContext()
		defer stop()

		q, err := queue.New(cfg)
		if err != nil {
			return err
		}
		defer q.Close()

		template, err := payload.Load(cfg.PayloadFile)
		if err != nil {
			return err
		}

		var mu sync.Mutex
		received := 0
		done := make(chan struct{})

		subCtx, subStop := context.WithCancel(ctx)
		defer subStop()

		subErr := make(chan error, 1)
		go func() {
			subErr <- q.Subscribe(subCtx, func(_ context.Context, msg queue.Message) error {
				fmt.Printf("recv  id=%s body=%s\n", msg.ID, msg.Body)
				mu.Lock()
				received++
				n := received
				mu.Unlock()
				if n >= cfg.Count {
					select {
					case <-done:
					default:
						close(done)
					}
				}
				return nil
			})
		}()

		// Give the subscriber time to create the group, queue, or topic.
		time.Sleep(2 * time.Second)

		fmt.Printf("demo: %d message(s) via %s\n", cfg.Count, describeTarget())
		for i := 1; i <= cfg.Count; i++ {
			msg := payload.NewMessage(template, i, time.Now())
			if err := q.Publish(ctx, msg); err != nil {
				return err
			}
			fmt.Printf("sent  seq=%d id=%s\n", i, msg.ID)

			if i < cfg.Count {
				select {
				case <-ctx.Done():
					return nil
				case <-time.After(cfg.Interval):
				}
			}
		}

		select {
		case <-done:
			fmt.Printf("demo: received all %d message(s)\n", cfg.Count)
		case err := <-subErr:
			if err != nil {
				return err
			}
			return fmt.Errorf("subscriber stopped before all messages arrived")
		case <-time.After(30 * time.Second):
			mu.Lock()
			n := received
			mu.Unlock()
			return fmt.Errorf("timeout: received %d of %d message(s)", n, cfg.Count)
		case <-ctx.Done():
			return nil
		}

		subStop()
		select {
		case <-subErr:
		case <-time.After(time.Second):
		}
		return nil
	},
}

// describeTarget names the broker and the destination for log lines.
func describeTarget() string {
	switch cfg.Broker {
	case "redis":
		return fmt.Sprintf("redis %s %q", cfg.Redis.Mode, cfg.Redis.Stream)
	case "rabbitmq":
		return fmt.Sprintf("rabbitmq queue %q", cfg.RabbitMQ.Queue)
	case "kafka":
		return fmt.Sprintf("kafka %s topic %q", cfg.Kafka.Flavor, cfg.Kafka.Topic)
	default:
		return cfg.Broker
	}
}

func init() {
	rootCmd.AddCommand(demoCmd)
}
