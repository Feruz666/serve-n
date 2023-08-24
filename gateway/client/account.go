package client

import (
	"fmt"

	"github.com/Feruz666/serve/gateway/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
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

func (c *AccountClient) ConnectToKafka() (*kafka.Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"acks":              "all",
	})

	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to kafka: %w", err)
	}

	return p, nil

}
