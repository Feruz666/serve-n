package client

import "github.com/confluentinc/confluent-kafka-go/kafka"

type KafkaProducer struct {
	producer   *kafka.Producer
	topic      string
	deliveryCh chan kafka.Event
}

func NewKafkaProducer(p *kafka.Producer, topic string) *KafkaProducer {
	return &KafkaProducer{
		producer:   p,
		topic:      topic,
		deliveryCh: make(chan kafka.Event, 1000),
	}
}
