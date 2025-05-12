package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

func main() {
	// Kafka broker configuration
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	// Create a new producer
	brokers := []string{"localhost:9092"}
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}
	defer producer.Close()

	// Create a signal channel to handle interruption
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Topic to send messages to
	topic := "test-topic"

	// Send messages until interrupted
	count := 1
	for {
		select {
		case <-signals:
			fmt.Println("\nInterrupted, stopping producer...")
			return
		default:
			message := fmt.Sprintf("Message %d", count)

			// Create a message
			msg := &sarama.ProducerMessage{
				Topic: topic,
				Value: sarama.StringEncoder(message),
			}

			// Send the message
			partition, offset, err := producer.SendMessage(msg)
			if err != nil {
				log.Printf("Failed to send message: %v\n", err)
				continue
			}

			log.Printf("Message sent successfully! Topic: %s, Partition: %d, Offset: %d, Message: %s\n",
				topic, partition, offset, message)

			count++
		}
	}
}
