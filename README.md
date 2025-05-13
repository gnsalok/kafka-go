# Kafka Go Example

This is a simple example of a Kafka producer and consumer implementation in Go.

## Prerequisites

1. Go 1.21 or later
2. Docker installed
3. `github.com/Shopify/sarama` package

## Setup Docker ( Kafka and Zookeeper )

1. Setup network
```bash
docker network create kafka-network
```

2. Start zookeeper

```bash
docker run -d --name zookeeper --network kafka-network -p 2181:2181 wurstmeister/zookeeper
```

3. Start Kafka with the correct configuration

```bash
docker run -d --name kafka --network kafka-network -p 9092:9092 -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092 -e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092 -e KAFKA_ADVERTISED_HOST_NAME=localhost wurstmeister/kafka
```

4. Verify that both containers are running

```bash
docker ps
```



## Running the Application

1. Install Go dependencies:
   ```bash
   go mod download
   ```

2. First, start the consumer in one terminal:
   ```bash
   go run consumer/main.go
   ```

3. Then, start the producer in another terminal:
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


## Summary for Docker, Kafka and Zookeeper

1. Created a Docker network for Kafka and Zookeeper to communicate:
   - Allows containers to communicate using container names as hostnames
   - Isolates Kafka traffic from other Docker networks

2. Started Zookeeper container:
   - Runs on port 2181
   - Provides coordination service for Kafka
   - Uses wurstmeister/zookeeper image

3. Started Kafka container with proper configuration:
   - Connected to Zookeeper on zookeeper:2181
   - Set up listeners on 0.0.0.0:9092 to accept all connections
   - Exposed port 9092 for external access
   - Set advertised host to localhost for client connections
   - Uses wurstmeister/kafka image

4. Verified containers are running:
   - Used docker ps to check container status
   - Confirmed both Kafka and Zookeeper are up

## Troubleshooting

If you need to restart the Kafka setup:

```bash
# Stop and remove existing containers
docker stop kafka zookeeper
docker rm kafka zookeeper
```

## Maintainer

[gnsalok](https://github.com/gnsalok)