package kafka

import (
	"context"
	"encoding/json"

	"transactions-service/internal/ports/out/dto"

	"github.com/segmentio/kafka-go"
)

func (p *KafkaProducer) CreateUserBalanceCredit(content dto.AccountCreditCreation) error {
	val, err := json.Marshal(content)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte("key"),
		Value: val,
	}

	err = p.producer.WriteMessages(context.Background(), msg)
	return err
}
