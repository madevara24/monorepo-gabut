package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"queue/internal/queue"
)

var receiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "Consume messages from the queue",
	RunE: func(_ *cobra.Command, _ []string) error {
		ctx, stop := commandContext()
		defer stop()

		q, err := queue.New(cfg)
		if err != nil {
			return err
		}
		defer q.Close()

		fmt.Printf("receive: listening on %s, press Ctrl+C to stop\n", describeTarget())
		return q.Subscribe(ctx, func(_ context.Context, msg queue.Message) error {
			fmt.Printf("recv  id=%s sentAt=%s body=%s\n",
				msg.ID, msg.SentAt.Format(time.RFC3339), msg.Body)
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(receiveCmd)
}
