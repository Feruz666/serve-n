package client

import "github.com/confluentinc/confluent-kafka-go/kafka"

type MessageSender struct {
	producer   *kafka.Producer
	topic      string
	deliveryCh chan kafka.Event
}

func NewMessageSender(p *kafka.Producer, topic string) *MessageSender {
	return &MessageSender{
		producer:   p,
		topic:      topic,
		deliveryCh: make(chan kafka.Event, 10000),
	}
}
