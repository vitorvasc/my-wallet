package kafka

import (
	"context"
	"encoding/json"
	"log"

	"account-balance-service/internal/core/services"

	"github.com/segmentio/kafka-go"
)

const (
	topic  = "users.account-balance-accreditations"
	broker = "kafka:9092"
)

func Consume(accountBalanceService *services.AccountBalanceService) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
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

		request := new(services.AccreditValueRequest)
		if err := json.Unmarshal(m.Value, request); err != nil {
			log.Printf("error unmarshalling message: %v", err)
			continue
		}

		// Accredit account balance
		if err := accountBalanceService.AccreditValue(request.UserID, request.Amount); err != nil {
			log.Printf("error accrediting account balance: %v", err)
			continue
		}
	}
}
