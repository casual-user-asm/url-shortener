package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func StartConsumer() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"kafka:9093"},
		Topic:    "shortened-urls",
		GroupID:  "shortened-urls-group",
		MinBytes: 1,
		MaxBytes: 10 * 1024 * 1024,
	})

	fmt.Println("Kafka Consumer Started")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading message: ", err)
			continue
		}
		fmt.Printf("Received: OriginalURL: %s => ShortURL: http://localhost:8080/%s\n", msg.Key, msg.Value)
	}
}
