package usecase

import (
	"context"
	"skeleton/domain/ddrivers/validation"
	"skeleton/lib/helper"
	"skeleton/pb/generic"
)

func (u *service) Delete(ctx context.Context, in *generic.Id) (*generic.BoolMessage, error) {
	select {
	case <-ctx.Done():
		return nil, helper.ContextError(ctx)
	default:
	}

	dValidation := validation.NewValidation(u.log, u.driverRepo)
	err := dValidation.Delete(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	err = u.driverRepo.Delete(ctx, in)
	if err != nil {
		return nil, err
	}

	return &generic.BoolMessage{IsTrue: true}, nil
}
