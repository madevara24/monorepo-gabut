package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"

	"queue/internal/config"
)

var cfg config.Config

var rootCmd = &cobra.Command{
	Use:   "queue",
	Short: "Learn how queues work with Redis, RabbitMQ, and Kafka",
	PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
		return resolve(cmd)
	},
}

// Execute runs the command line interface.
func Execute() {
	_ = godotenv.Load() // .env is optional
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func init() {
	p := rootCmd.PersistentFlags()

	p.String("broker", "", "queue backend: redis | rabbitmq | kafka")
	p.String("payload", "", "path to the json payload template")
	p.Int("count", 0, "number of messages to send")
	p.Duration("interval", 0, "delay between messages, for example 1s")

	p.String("redis-addr", "", "redis address")
	p.String("redis-password", "", "redis password")
	p.Int("redis-db", 0, "redis database number")
	p.String("redis-mode", "", "redis mode: streams | list")
	p.String("redis-stream", "", "redis stream or list key")
	p.String("redis-group", "", "redis consumer group")

	p.String("rabbit-url", "", "rabbitmq amqp url")
	p.String("rabbit-queue", "", "rabbitmq queue name")

	p.String("kafka-flavor", "", "kafka flavor: kafka | redpanda")
	p.StringSlice("kafka-brokers", nil, "kafka broker addresses")
	p.String("kafka-topic", "", "kafka topic")
	p.String("kafka-group", "", "kafka consumer group")
}
