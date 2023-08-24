package client

import (
	"context"
	"fmt"

	pb "github.com/Feruz666/serve/gateway/protos/account/proto"
)

func (c *AccountClient) RegExp(ctx context.Context, username string) (*pb.RegExpResponse, error) {
	p, err := c.ConnectToKafka()
	if err != nil {
		return nil, fmt.Errorf("c.ConnectToKafka failed %w", err)
	}

	topic := "account-topic"

	if err != nil {
		return nil, fmt.Errorf("c.RegExp failed %w", err)
	}

	return nil, nil
}
