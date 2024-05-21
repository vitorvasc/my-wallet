package kafka

import (
	"account-balance-service/internal/core/services"
	in "account-balance-service/internal/ports/in/kafka"
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

const (
	topic  = "users.account-balance-accreditations"
	broker = "KAFKA_BROKER"
)

func Consume(service services.AccountBalance) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{os.Getenv(broker)},
		Topic:   topic,
		GroupID: "account-balance-service",
	})
	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("error reading message: %v", err)
			continue
		}

		request := new(in.AccreditationRequest)
		if err = json.Unmarshal(m.Value, request); err != nil {
			log.Printf("error unmarshalling message: %v", err)
			continue
		}

		// Accredit account balance
		if err = service.AccreditValue(request.UserID, request.Amount); err != nil {
			log.Printf("error accrediting account balance: %v", err)
			continue
		}
	}
}
