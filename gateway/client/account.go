package client

import (
	"fmt"
	"log"

	"github.com/Feruz666/serve/gateway/config"
	account "github.com/Feruz666/serve/gateway/protos/account/protos"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/protobuf/proto"
)

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

// AccountClient client for account service
type AccountClient struct {
	config *config.Config
}

func NewAccountClient(config *config.Config) *AccountClient {
	return &AccountClient{
		config: config,
	}
}

func (c *AccountClient) CreatePerson() error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "foo",
		"acks":              "all",
	})

	if err != nil {
		return fmt.Errorf("kafka.NewProducer failed - %w", err)
	}

	topic := "account-topic"
	pers := &account.Person{
		Name:  "Feruz",
		Id:    223,
		Email: "asd@gmail.com",
	}

	serializedPers, err := proto.Marshal(pers)

	ms := NewMessageSender(p, "account-topic")

	for i := 0; i < 10; i++ {
		err := ms.producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: serializedPers,
		}, ms.deliveryCh)

		if err != nil {
			log.Fatal(err)
		}

		e := <-ms.deliveryCh
		m := e.(*kafka.Message)
		if m.TopicPartition.Error != nil {
			fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
		} else {
			fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
				*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		}

		
	}

	return nil
}
