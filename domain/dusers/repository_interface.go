package dusers

import (
	"context"
	"ksuser/pb/users"
)

type UserRepoInterface interface {
	Create(ctx context.Context) error
	GetPb() *users.User
	SetPb(*users.User)
}
