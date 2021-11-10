package dusers

import (
	"context"
	"ksuser/pb/users"
)

type UserUsecaseInterface interface {
	Create(ctx context.Context, in *users.User) (*users.User, error)
}
