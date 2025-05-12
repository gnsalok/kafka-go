# kafka-go

This is a simple example of a Kafka producer and consumer implementation in Go.

## Prerequisites

1. Go 1.21 or later
2. Apache Kafka running locally (or update broker addresses in code)
3. `github.com/Shopify/sarama` package

## Setup

1. Start Kafka and Zookeeper locally
2. Install dependencies:
   ```bash
   go mod download
   ```

## Running the Application

1. First, start the consumer in one terminal:
   ```bash
   go run consumer/main.go
   ```

2. Then, start the producer in another terminal:
   ```bash
   go run producer/main.go
   ```

## Testing

1. The producer will continuously send numbered messages to the "test-topic" topic
2. The consumer will read these messages and print them to the console
3. You can stop either program by pressing Ctrl+C

## Expected Output

Producer output will look like:
```
Message sent successfully! Topic: test-topic, Partition: 0, Offset: 0, Message: Message 1
Message sent successfully! Topic: test-topic, Partition: 0, Offset: 1, Message: Message 2
...
```

Consumer output will look like:
```
Message received: Topic=test-topic, Partition=0, Offset=0, Value=Message 1
Message received: Topic=test-topic, Partition=0, Offset=1, Value=Message 2
...
```

## Docker

To run Kafka using Docker:
```bash
This error occurs because Kafka is not running or is not accessible at localhost:9092. Let me help you fix this:

1. First, let's make sure Kafka is running. I'll provide a Docker command to start both Kafka and Zookeeper:

```bash
docker run -d --name kafka-broker -p 2181:2181 -p 9092:9092 -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092 -e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092 -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 -e KAFKA_BROKER_ID=1 -e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 -e KAFKA_ZOOKEEPER_CONNECT=localhost:2181 wurstmeister/kafka:latest

``` 