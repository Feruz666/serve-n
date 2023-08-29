package kafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

type KafkaProducerConfig struct {
	Topic            string
	BootstrapServers []string
	Acks             int
	Clientid         string
	DeliveryCh       chan kafka.Event
}

func NewKafkaProducer(topic string, clientid string, acks int, bootstrapServers []string) *KafkaProducerConfig {
	return &KafkaProducerConfig{
		Topic:            topic,
		BootstrapServers: bootstrapServers,
		Acks:             acks,
		Clientid:         clientid,
		DeliveryCh:       make(chan kafka.Event, 1000),
	}
}
