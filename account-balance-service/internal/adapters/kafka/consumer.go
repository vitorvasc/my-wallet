package kafka

import (
	"account-balance-service/internal/adapters/config"
	"account-balance-service/internal/core/services"
	in "account-balance-service/internal/ports/in/kafka"
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

const (
	topic  = "users.account-balance-accreditations"
	broker = "kafka:9092"
)

func Consume() {
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

		request := new(in.AccreditationRequest)
		if err = json.Unmarshal(m.Value, request); err != nil {
			log.Printf("error unmarshalling message: %v", err)
			continue
		}

		// Accredit account balance
		accountBalanceService := config.Container.Get(config.AccountBalanceService).(services.AccountBalance)
		if err = accountBalanceService.AccreditValue(request.UserID, request.Amount); err != nil {
			log.Printf("error accrediting account balance: %v", err)
			continue
		}
	}
}
