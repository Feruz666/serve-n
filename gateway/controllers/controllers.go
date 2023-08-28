package controllers

import "github.com/Feruz666/serve-n/gateway/controllers/interfaces"


// Controllers for account gateway
type Controllers struct {
	client interfaces.AccountClient
}

func NewControllers(client interfaces.AccountClient) *Controllers {
	return &Controllers{
		client: client,
	}
}
