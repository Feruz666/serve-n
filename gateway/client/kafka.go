package client

import "github.com/confluentinc/confluent-kafka-go/kafka"

type MessageSender struct {
	Producer   *kafka.Producer
	Topic      string
	DeliveryCh chan kafka.Event
}

func NewMessageSender(p *kafka.Producer, topic string) *MessageSender {
	return &MessageSender{
		Producer:   p,
		Topic:      topic,
		DeliveryCh: make(chan kafka.Event, 10000),
	}
}
