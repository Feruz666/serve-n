package client

import (
	"fmt"
	"log"
	"time"

	"github.com/Feruz666/serve-n/gateway/config"
	account "github.com/Feruz666/serve-n/gateway/protos/account/protos"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/protobuf/proto"
)

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
		err := ms.Producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: int32(kafka.PartitionAny),
			},
			Value: serializedPers,
		}, ms.DeliveryCh)

		if err != nil {
			log.Fatal(err)
		}

		e := <-ms.DeliveryCh
		m := e.(*kafka.Message)
		if m.TopicPartition.Error != nil {
			fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
		} else {
			fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
				*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		}

		time.Sleep(time.Second * 2)
	}

	return nil
}
