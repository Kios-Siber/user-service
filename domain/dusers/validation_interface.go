package dusers

import (
	"context"
	"ksuser/pb/users"
)

type UserValidationInterface interface {
	Create(ctx context.Context, id *users.User) error
}
