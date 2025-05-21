package messages

import (
	"net"
	"strconv"

	"github.com/segmentio/kafka-go"
)

func CreateTopic(brokerAddress string, topic string) error {
	conn, err := kafka.Dial("tcp", brokerAddress)
	if err != nil {
		return err
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return err
	}

	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		return err
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{{Topic: topic, NumPartitions: 1, ReplicationFactor: 1}}
	err = controllerConn.CreateTopics(topicConfigs...)
	return err
}

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
