# Queue Learning Service

This is a small Go program that shows how a message queue works. The program is one binary, and the binary can send and receive messages. It works with three brokers: Redis, RabbitMQ, and Kafka. A broker is a server that stores and moves messages.

## Concepts

The words below appear in the code and in the logs.

- A producer, or sender, puts a message into the queue.
- A consumer, or receiver, takes a message out of the queue.
- An acknowledgement, or ack, tells the broker that the consumer finished the message.
- Redelivery happens when the broker gives an unacknowledged message to a consumer again.
- A consumer group is a set of consumers that share the work of one queue.
- An offset is the position of one message inside a Kafka topic.

## Requirements

- Go 1.23 or newer
- Docker and Docker Compose

## First setup

1. Run `make tidy` to download the Go dependencies.
2. Run `make build` to build the binary into `bin/queue`.

## Run a demo

The demo command starts one consumer and one producer in the same process. The consumer prints every message that it receives.

Redis with Streams:

1. Run `make up-redis` to start Redis.
2. Run `make demo` to send and receive five messages.
3. Run `make down` to stop Redis and delete its data.

Redis with a list:

1. Run `make up-redis`.
2. Run `make demo REDIS_MODE=list REDIS_STREAM=demo-list`.
3. Run `make down`.

Redis stores a stream and a list in different ways, and one key holds one type. Use a different key when you switch between `streams` and `list` on a live Redis. The make variable `REDIS_STREAM` sets the key. `make down` also deletes the data, so a fresh Redis accepts the default key `demo`.

RabbitMQ:

1. Run `make up-rabbitmq`.
2. Run `make demo BROKER=rabbitmq`.
3. Run `make down`.

Kafka:

1. Run `make up-kafka`.
2. Run `make demo BROKER=kafka`.
3. Run `make down`.

Redpanda:

1. Run `make up-redpanda`.
2. Run `make demo BROKER=kafka KAFKA_FLAVOR=redpanda`.
3. Run `make down`.

## Send and receive in two terminals

Two terminals show the producer and the consumer as separate processes. The example below uses Redis.

1. Run `make up-redis` in the first terminal.
2. Run `make receive` in the first terminal. The consumer waits for messages.
3. Run `make send` in the second terminal. The producer publishes five messages.
4. Look at the first terminal. The consumer prints the five messages.
5. Press Ctrl+C in the first terminal to stop the consumer.
6. Run `make down`.

## Make commands

Setup:

- `make tidy` downloads the Go dependencies.
- `make build` builds the binary into `bin/queue`.
- `make run ARGS="demo"` runs the binary with your own arguments.

Brokers:

- `make up-redis` starts Redis on port 6379.
- `make up-rabbitmq` starts RabbitMQ on port 5672. The management page is on port 15672, user `guest`, password `guest`.
- `make up-kafka` starts Kafka in KRaft mode on port 9092.
- `make up-redpanda` starts Redpanda on port 19092.
- `make up` starts the broker in the `BROKER` variable.
- `make down` stops every broker and deletes its data.
- `make ps` shows the running containers.
- `make logs` follows the broker logs.

Run:

- `make send` publishes messages.
- `make receive` consumes messages.
- `make demo` publishes and consumes in one process.

## How each broker works

Redis has two modes.

Streams mode uses a consumer group. Every message has an entry ID. The consumer reads with `XREADGROUP` and confirms with `XACK`. A message stays in the stream until the consumer confirms it. If the consumer stops before the ack, the message returns later. The queue keeps a history.

List mode uses `LPUSH` and `BRPOP`. The push adds a message to the head of a list. The pop removes a message from the tail. The pop deletes the message at once. A consumer that stops after the pop loses that message, and no other consumer sees it. This mode has no ack and no redelivery.

RabbitMQ uses one durable queue. The producer publishes to the default exchange with the queue name as the routing key. The consumer confirms a message with `Ack`. A handler error sends the message back with `Nack` and `requeue=true`. RabbitMQ then delivers the message again. The management page shows the queue depth and the delivery rate.

Kafka uses a topic with partitions. A topic is an append-only log. The producer adds a message to the end of the log. A consumer group tracks an offset for each partition. The consumer confirms a message by committing the offset. Kafka keeps the message in the log after the commit, and the retention time decides when Kafka deletes it. A handler error leaves the offset uncommitted, so the message returns on the next read.

Redpanda speaks the Kafka API. The Go client uses the same code for Kafka and Redpanda. Only the address changes.

## The payload template

The file `message.json` holds the payload. Edit this file to change the message body. The sender copies the file for every message. Two placeholders change per message:

- `{{seq}}` becomes the sequence number.
- `{{time}}` becomes the send time in RFC 3339 format.

The sender also adds two fields to the envelope: `id` and `sentAt`. The `id` is a UUID that is unique for every message.

## Configuration

A command line flag wins over an environment variable. An environment variable wins over a default value. Copy `.env.example` to `.env` to set the values in one place.

General:

- `--broker` or `BROKER` selects the broker. Values: `redis`, `rabbitmq`, `kafka`. Default `redis`.
- `--payload` or `PAYLOAD_FILE` sets the payload template path. Default `message.json`.
- `--count` or `COUNT` sets the number of messages that `send` and `demo` publish. Default `5`.
- `--interval` or `INTERVAL` sets the delay between messages, for example `1s`. Default `1s`.

Redis:

- `--redis-addr` or `REDIS_ADDR`. Default `localhost:6379`.
- `--redis-password` or `REDIS_PASSWORD`. Default empty.
- `--redis-db` or `REDIS_DB`. Default `0`.
- `--redis-mode` or `REDIS_MODE`. Values: `streams`, `list`. Default `streams`.
- `--redis-stream` or `REDIS_STREAM` sets the stream name or the list key. Default `demo`. Use a different key for `streams` and `list` on one live Redis.
- `--redis-group` or `REDIS_GROUP` sets the consumer group for streams mode. Default `demo-group`.

RabbitMQ:

- `--rabbit-url` or `RABBIT_URL`. Default `amqp://guest:guest@localhost:5672/`.
- `--rabbit-queue` or `RABBIT_QUEUE`. Default `demo`.

Kafka:

- `--kafka-flavor` or `KAFKA_FLAVOR` selects the default broker address. Values: `kafka`, `redpanda`. Default `kafka`.
- `--kafka-brokers` or `KAFKA_BROKERS` sets the addresses, separated by commas. Default from the flavor: `localhost:9092` for Kafka, `localhost:19092` for Redpanda.
- `--kafka-topic` or `KAFKA_TOPIC`. Default `demo`.
- `--kafka-group` or `KAFKA_GROUP`. Default `demo-group`.

## Project layout

```
main.go                     program entry point
cmd/                        command line commands: send, receive, demo
internal/config/            the configuration struct
internal/queue/             the Queue interface and three implementations
internal/payload/           payload template loading and placeholder replacement
message.json                editable payload template
docker-compose.yml          one profile for each broker
Makefile                    all commands
```

To add a new broker, create one file in `internal/queue`, add the implementation to the `New` function in `internal/queue/queue.go`, and add the options to `cmd/`.
