package messages

import (
	"github.com/segmentio/kafka-go"
)

func ConnectKafkaWriter(brokerAddress string, topic string) *kafka.Writer {
	writer := &kafka.Writer{
		Addr:         kafka.TCP(brokerAddress),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
		Async:        false,
	}
	return writer
}

func ConnectKafkaReader(groupID string, topic string) *kafka.Reader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092", "localhost:9093"},
		Topic:   topic,
		GroupID: groupID,
	})

	return reader
}
