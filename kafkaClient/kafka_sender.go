package kafkaclient

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type Sender struct {
	Writer *kafka.Writer
}

func (s *Sender) Send(topic string, m []byte) error {
	km := kafka.Message{
		Topic: topic,
		Key:   []byte(uuid.New().String()),
		Value: m,
		Time:  time.Time{},
	}
	return s.Writer.WriteMessages(context.Background(), km)
}

func NewSender(conf *Conf) *Sender {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  conf.ProducerBootstrapServers,
		Balancer: &kafka.Hash{},
	})
	return &Sender{Writer: w}
}
