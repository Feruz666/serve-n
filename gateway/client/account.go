package client

import (
	"context"
	"fmt"

	"github.com/Feruz666/serve-n/gateway/config"
	kfk "github.com/Feruz666/serve-n/gateway/controllers/kafka"
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

func (c *AccountClient) CreatePerson(ctx context.Context, user *account.Person) error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "1",
		"acks":              "all",
	})

	pr := kfk.NewKafkaProducerConfig(p, kfk.AccountTopic, "sender", "gateway")

	if err != nil {
		return fmt.Errorf("kafka.NewProducer failed - %w", err)
	}

	serializedPers, err := proto.Marshal(user)
	if err != nil {
		return fmt.Errorf("proto.Marshal() failed - %w", err)
	}

	if err := pr.SendMessage(serializedPers); err != nil {
		return fmt.Errorf("pr.SendMessage() failed - %w", err)
	}

	fmt.Println("We are here!!!")

	return nil
}
