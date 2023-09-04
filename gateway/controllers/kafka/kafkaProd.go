package kafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducerConfig struct {
	Producer    *kafka.Producer
	Topic       string
	HeaderKey   string
	HeaderValue string
	DeliveryCh  chan kafka.Event
}

// NewKafkaProducer creates a new kafka producer ...
func NewKafkaProducerConfig(producer *kafka.Producer, topic, headerKey, headerValue string) *KafkaProducerConfig {
	return &KafkaProducerConfig{
		Producer:    producer,
		Topic:       topic,
		HeaderKey:   headerKey,
		HeaderValue: headerValue,
		DeliveryCh:  make(chan kafka.Event, 1000),
	}
}

func (kp *KafkaProducerConfig) SendMessage(msg []byte) error {
	headers := []kafka.Header{{Key: kp.HeaderKey, Value: []byte(kp.HeaderValue)}}

	err := kp.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &kp.Topic,
			Partition: int32(kafka.PartitionAny),
		},
		Value:   msg,
		Headers: headers,
	}, kp.DeliveryCh)

	if err != nil {
		log.Fatal(err)
	}

	<-kp.DeliveryCh

	return nil
}
