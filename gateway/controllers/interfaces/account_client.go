package interfaces

import (
	"context"

	account "github.com/Feruz666/serve-n/gateway/protos/account/protos"
)

type AccountClient interface {
	CreatePerson(ctx context.Context, user *account.Person) error
}
