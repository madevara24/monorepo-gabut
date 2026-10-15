package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"queue/internal/payload"
	"queue/internal/queue"
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Publish messages to the queue",
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

		fmt.Printf("send: %d message(s) to %s\n", cfg.Count, describeTarget())
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
		return nil
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
}
