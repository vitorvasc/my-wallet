package kafka

import (
	"github.com/segmentio/kafka-go"
)

const (
	address = "kafka:9092"
	topic   = "users.account-balance-accreditations"
)

type KafkaProducer struct {
	producer *kafka.Writer
}

func NewKafkaProducer() *KafkaProducer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(address),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	return &KafkaProducer{producer: writer}
}

func (p *KafkaProducer) ProducerClose() {
	p.producer.Close()
}
