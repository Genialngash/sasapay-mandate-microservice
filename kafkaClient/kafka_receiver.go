package kafkaclient

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type HandleFunc func([]byte)

type Receiver struct {
	Reader *kafka.Reader
	Handle HandleFunc
}

func (r *Receiver) ReceiveAndHandle() {
	go func() {
		for {
			m, err := r.Reader.ReadMessage(context.Background())
			if err != nil {
				log.Println(err)
				continue
			}
			if m.Value == nil {
				log.Println(fmt.Errorf("topic: %s received empty record %s", m.Topic, m.Key))
				continue
			}
			go r.Handle(m.Value)
		}
	}()
}

func NewReceiver(conf *Conf, topics []string, h HandleFunc) *Receiver {
	nr := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        conf.ConsumerBootstrapServers,
		GroupID:        conf.GroupId,
		StartOffset:    kafka.LastOffset,
		GroupTopics:    topics,
		ReadBackoffMin: 1,
		ReadBackoffMax: 10,
		MinBytes:       1,    // 1B
		MaxBytes:       10e3, // 10MB
	})
	r := &Receiver{Reader: nr, Handle: h}
	return r
}
