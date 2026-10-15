package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

// resolve builds the config from three sources.
// A flag wins over an environment variable. An environment variable wins over a default.
func resolve(cmd *cobra.Command) error {
	cfg.Broker = pickString(cmd, "broker", "BROKER", "redis")
	cfg.PayloadFile = pickString(cmd, "payload", "PAYLOAD_FILE", "message.json")
	cfg.Count = pickInt(cmd, "count", "COUNT", 5)
	cfg.Interval = pickDuration(cmd, "interval", "INTERVAL", time.Second)

	cfg.Redis.Addr = pickString(cmd, "redis-addr", "REDIS_ADDR", "localhost:6379")
	cfg.Redis.Password = pickString(cmd, "redis-password", "REDIS_PASSWORD", "")
	cfg.Redis.DB = pickInt(cmd, "redis-db", "REDIS_DB", 0)
	cfg.Redis.Mode = pickString(cmd, "redis-mode", "REDIS_MODE", "streams")
	cfg.Redis.Stream = pickString(cmd, "redis-stream", "REDIS_STREAM", "demo")
	cfg.Redis.Group = pickString(cmd, "redis-group", "REDIS_GROUP", "demo-group")

	cfg.RabbitMQ.URL = pickString(cmd, "rabbit-url", "RABBIT_URL", "amqp://guest:guest@localhost:5672/")
	cfg.RabbitMQ.Queue = pickString(cmd, "rabbit-queue", "RABBIT_QUEUE", "demo")

	cfg.Kafka.Flavor = pickString(cmd, "kafka-flavor", "KAFKA_FLAVOR", "kafka")
	cfg.Kafka.Brokers = pickStringSlice(cmd, "kafka-brokers", "KAFKA_BROKERS", defaultKafkaBrokers(cfg.Kafka.Flavor))
	cfg.Kafka.Topic = pickString(cmd, "kafka-topic", "KAFKA_TOPIC", "demo")
	cfg.Kafka.Group = pickString(cmd, "kafka-group", "KAFKA_GROUP", "demo-group")

	return validate()
}

func validate() error {
	switch cfg.Broker {
	case "redis", "rabbitmq", "kafka":
	default:
		return fmt.Errorf("unknown broker %q, use redis, rabbitmq, or kafka", cfg.Broker)
	}

	if cfg.Broker == "redis" {
		switch cfg.Redis.Mode {
		case "streams", "list":
		default:
			return fmt.Errorf("unknown redis mode %q, use streams or list", cfg.Redis.Mode)
		}
	}

	if cfg.Broker == "kafka" {
		switch cfg.Kafka.Flavor {
		case "kafka", "redpanda":
		default:
			return fmt.Errorf("unknown kafka flavor %q, use kafka or redpanda", cfg.Kafka.Flavor)
		}
	}

	if cfg.Count < 0 {
		return fmt.Errorf("count must be 0 or more")
	}
	return nil
}

func defaultKafkaBrokers(flavor string) []string {
	if flavor == "redpanda" {
		return []string{"localhost:19092"}
	}
	return []string{"localhost:9092"}
}

// commandContext stops the command on Ctrl+C or on a termination signal.
func commandContext() (context.Context, context.CancelFunc) {
	return signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
}

func pickString(cmd *cobra.Command, flag, env, fallback string) string {
	if f := cmd.Flags().Lookup(flag); f != nil && f.Changed {
		value, _ := cmd.Flags().GetString(flag)
		return value
	}
	if value, ok := os.LookupEnv(env); ok && value != "" {
		return value
	}
	return fallback
}

func pickInt(cmd *cobra.Command, flag, env string, fallback int) int {
	if f := cmd.Flags().Lookup(flag); f != nil && f.Changed {
		value, _ := cmd.Flags().GetInt(flag)
		return value
	}
	if value, ok := os.LookupEnv(env); ok && value != "" {
		if n, err := strconv.Atoi(value); err == nil {
			return n
		}
	}
	return fallback
}

func pickDuration(cmd *cobra.Command, flag, env string, fallback time.Duration) time.Duration {
	if f := cmd.Flags().Lookup(flag); f != nil && f.Changed {
		value, _ := cmd.Flags().GetDuration(flag)
		return value
	}
	if value, ok := os.LookupEnv(env); ok && value != "" {
		if d, err := time.ParseDuration(value); err == nil {
			return d
		}
	}
	return fallback
}

func pickStringSlice(cmd *cobra.Command, flag, env string, fallback []string) []string {
	if f := cmd.Flags().Lookup(flag); f != nil && f.Changed {
		value, _ := cmd.Flags().GetStringSlice(flag)
		return value
	}
	if value, ok := os.LookupEnv(env); ok && value != "" {
		var out []string
		for _, part := range strings.Split(value, ",") {
			if part = strings.TrimSpace(part); part != "" {
				out = append(out, part)
			}
		}
		return out
	}
	return fallback
}
