# Broker selection. Set BROKER on the command line: make demo BROKER=rabbitmq
BROKER ?= redis
REDIS_MODE ?= streams
REDIS_STREAM ?= demo
KAFKA_FLAVOR ?= kafka
BIN := bin/queue
COMPOSE := docker compose

.PHONY: help tidy build run send receive demo \
        up up-redis up-rabbitmq up-kafka up-redpanda down ps logs

help:
	@echo "Queue learning service"
	@echo ""
	@echo "Setup:"
	@echo "  make tidy                 Download Go dependencies"
	@echo "  make build                Build the binary into $(BIN)"
	@echo ""
	@echo "Start one broker:"
	@echo "  make up-redis             Start Redis"
	@echo "  make up-rabbitmq          Start RabbitMQ"
	@echo "  make up-kafka             Start Kafka (KRaft)"
	@echo "  make up-redpanda          Start Redpanda"
	@echo "  make up                   Start the broker in BROKER ($(BROKER))"
	@echo "  make down                 Stop all brokers and delete their data"
	@echo ""
	@echo "Run the service:"
	@echo "  make send                 Publish messages"
	@echo "  make receive              Consume messages"
	@echo "  make demo                 Publish and consume in one process"
	@echo ""
	@echo "Options:"
	@echo "  BROKER=redis|rabbitmq|kafka    (current: $(BROKER))"
	@echo "  REDIS_MODE=streams|list        (current: $(REDIS_MODE))"
	@echo "  REDIS_STREAM=<key>             (current: $(REDIS_STREAM))"
	@echo "  KAFKA_FLAVOR=kafka|redpanda    (current: $(KAFKA_FLAVOR))"
	@echo ""
	@echo "Examples:"
	@echo "  make up-redis && make demo"
	@echo "  make up-redis && make demo REDIS_MODE=list REDIS_STREAM=demo-list"
	@echo "  make up-rabbitmq && make demo BROKER=rabbitmq"
	@echo "  make up-kafka && make demo BROKER=kafka"
	@echo "  make up-redpanda && make demo BROKER=kafka KAFKA_FLAVOR=redpanda"

tidy:
	go mod tidy

build:
	go build -o $(BIN) .

run:
	go run . $(ARGS)

send:
	go run . --broker $(BROKER) --redis-mode $(REDIS_MODE) --redis-stream $(REDIS_STREAM) --kafka-flavor $(KAFKA_FLAVOR) send

receive:
	go run . --broker $(BROKER) --redis-mode $(REDIS_MODE) --redis-stream $(REDIS_STREAM) --kafka-flavor $(KAFKA_FLAVOR) receive

demo:
	go run . --broker $(BROKER) --redis-mode $(REDIS_MODE) --redis-stream $(REDIS_STREAM) --kafka-flavor $(KAFKA_FLAVOR) demo

up:
	$(COMPOSE) --profile $(BROKER) up -d

up-redis:
	$(COMPOSE) --profile redis up -d

up-rabbitmq:
	$(COMPOSE) --profile rabbitmq up -d

up-kafka:
	$(COMPOSE) --profile kafka up -d

up-redpanda:
	$(COMPOSE) --profile redpanda up -d

down:
	$(COMPOSE) --profile redis --profile rabbitmq --profile kafka --profile redpanda down -v

ps:
	$(COMPOSE) ps

logs:
	$(COMPOSE) logs -f
