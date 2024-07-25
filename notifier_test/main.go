package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type PaymentReceived struct {
	ChargeID  string    `json:"charge_id"`
	PartID    string    `json:"part_id"`
	Amount    float64   `json:"amount"`
	TotalPaid float64   `json:"total_paid"`
	PaidBy    string    `json:"paid_by"`
	PaidAt    time.Time `json:"paid_at"`
}

func main() {
	ctx := context.Background()

	// make a writer that produces to topic-A, using the least-bytes distribution
	writer := &kafka.Writer{
		Addr:     kafka.TCP("kafka:9092", "kafka:9093", "kafka:9094"),
		Balancer: &kafka.RoundRobin{},
	}

	// msg := PaymentReceived{
	// 	ChargeID:  "ch_1J1J9v2eZvKYlo2C5",
	// 	PartID:    "pt_1J1J9v2eZvKYlo2C6",
	// 	Amount:    100.00,
	// 	TotalPaid: 100.00,
	// 	PaidBy:    "acct_1J1J9v2eZvKYlo2C7",
	// 	PaidAt:    time.Now(),
	// }

	// msgValue, err := json.Marshal(msg)
	// if err != nil {
	// 	log.Fatal("failed to marshal message:", err)
	// }

	err := writer.WriteMessages(
		ctx,
		kafka.Message{
			Topic: "topic-A",
			// Topic: "payments.charge.payment-received",
			// Key:   []byte("Key-A"),
			Value: []byte("Hello Guilherme!"),
			// Value: msgValue,
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := writer.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
