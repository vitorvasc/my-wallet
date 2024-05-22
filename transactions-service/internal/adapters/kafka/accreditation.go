package kafka

import (
	"context"
	"encoding/json"
	out "transactions-service/internal/ports/out/repository"

	"github.com/segmentio/kafka-go"
)

func (p *KafkaProducer) AccreditateUserBalance(content out.UserBalanceAccreditation) error {
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
