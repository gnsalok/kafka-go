package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

func main() {
	// Kafka broker configuration
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// Create a new consumer
	brokers := []string{"localhost:9092"} // broker is the servers that hold the queue
	topic := "test-topic"
	group := "test-group"

	consumer, err := sarama.NewConsumerGroup(brokers, group, config)
	if err != nil {
		log.Fatalf("Error creating consumer group: %v", err)
	}
	defer consumer.Close()

	// Create a signal channel to handle interruption
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Create context that will be canceled on interrupt
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-signals
		cancel()
	}()

	// Start consuming messages
	handler := NewConsumerGroupHandler()
	topics := []string{topic}

	log.Printf("Starting to consume messages from topic: %s\n", topic)
	for {
		err := consumer.Consume(ctx, topics, handler)
		if err != nil {
			if ctx.Err() != nil {
				// Context was canceled
				fmt.Println("\nInterrupted, stopping consumer...")
				return
			}
			log.Printf("Error from consumer: %v\n", err)
		}
	}
}

// ConsumerGroupHandler represents the consumer group handler
type ConsumerGroupHandler struct {
	ready chan bool
}

// NewConsumerGroupHandler creates a new ConsumerGroupHandler
func NewConsumerGroupHandler() *ConsumerGroupHandler {
	return &ConsumerGroupHandler{
		ready: make(chan bool),
	}
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (h *ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	close(h.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (h *ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (h *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Message received: Topic=%s, Partition=%d, Offset=%d, Key=%s, Value=%s\n",
			message.Topic, message.Partition, message.Offset, string(message.Key), string(message.Value))
		session.MarkMessage(message, "")
	}
	return nil
}
