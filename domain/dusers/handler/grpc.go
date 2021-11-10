package handler

import (
	"context"
	"ksuser/domain/dusers"
	"ksuser/pb/users"
)

type UserHandler struct {
	usecase dusers.UserUsecaseInterface
}

func NewUserHandler(usecase dusers.UserUsecaseInterface) *UserHandler {
	handler := new(UserHandler)
	handler.usecase = usecase
	return handler
}

func (u *UserHandler) Create(ctx context.Context, in *users.User) (*users.User, error) {
	return u.usecase.Create(ctx, in)
}
