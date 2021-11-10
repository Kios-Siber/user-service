package usecase

import (
	"context"
	"ksuser/domain/dusers/validation"
	"ksuser/lib/helper"
	"ksuser/pb/users"
)

func (u *userservice) Create(ctx context.Context, in *users.User) (*users.User, error) {
	select {
	case <-ctx.Done():
		return nil, helper.ContextError(ctx)
	default:
	}

	ctx, err := helper.GetMetadata(ctx)
	if err != nil {
		return nil, err
	}

	dValidation := validation.NewUserValidation(u.log, u.userRepo)
	err = dValidation.Create(ctx, in)
	if err != nil {
		return nil, err
	}

	u.userRepo.SetPb(in)

	err = u.userRepo.Create(ctx)
	if err != nil {
		return nil, err
	}

	return u.userRepo.GetPb(), nil
}
