package config

import "time"

// Config holds every value that a command needs.
// A command builds it from flags, environment variables, and default values.
type Config struct {
	Broker      string
	PayloadFile string
	Count       int
	Interval    time.Duration

	Redis    Redis
	RabbitMQ RabbitMQ
	Kafka    Kafka
}

type Redis struct {
	Addr     string
	Password string
	DB       int
	Mode     string // streams or list
	Stream   string
	Group    string
}

type RabbitMQ struct {
	URL   string
	Queue string
}

type Kafka struct {
	Flavor  string // kafka or redpanda
	Brokers []string
	Topic   string
	Group   string
}
