package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	ctx := context.Background()

	// make a new reader that consumes from topic-A, partition 0, at offset 42
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092", "kafka:9093", "kafka:9094"},
		Topic:   "topic-A",
		// Partition: 0,
		GroupID:     "my-group",
		MaxBytes:    10e6, // 10MB
		StartOffset: kafka.LastOffset,
	})
	// reader.SetOffset(2)

	for {
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			break
		}

		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := reader.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
