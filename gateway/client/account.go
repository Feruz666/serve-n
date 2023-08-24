package client

import (
	"fmt"

	"github.com/Feruz666/serve/gateway/config"
	account "github.com/Feruz666/serve/gateway/protos/account/protos"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/protobuf/proto"
)

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

// AccountClient client for account service
type AccountClient struct {
	config *config.Config
}

func NewAccountClient(config *config.Config) *AccountClient {
	return &AccountClient{
		config: config,
	}
}

func (c *AccountClient) ConnectToKafka() (*kafka.Producer, error) {
	p := NewKafkaProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"acks":              "all",
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to kafka: %w", err)
	}

	return p, nil
}

func (c AccountClient) CreatePerson() error {
	p, err := c.ConnectToKafka()

	if err != nil {
		fmt.Errorf("c.ConnectToKafka failed: %w", err)
	}

	topic := "account-topic"
	pers := &account.Person{
		Name:  "asd",
		Id:    223,
		Email: "asd@gmail.com",
	}

	buff, err := proto.Marshal(pers)

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: int32(kafka.PartitionAny),
		},
		Value: buff,
	},
		p.deliveryCh,
	)

	return nil
}
