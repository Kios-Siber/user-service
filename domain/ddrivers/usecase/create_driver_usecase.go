package usecase

import (
	"context"
	"ksuser/domain/ddrivers/validation"
	"ksuser/lib/helper"
	"ksuser/pb/drivers"
)

func (u *service) Create(ctx context.Context, in *drivers.Driver) (*drivers.Driver, error) {
	select {
	case <-ctx.Done():
		return nil, helper.ContextError(ctx)
	default:
	}

	dValidation := validation.NewValidation(u.log, u.driverRepo)
	err := dValidation.Create(ctx, in)
	if err != nil {
		return nil, err
	}

	u.driverRepo.SetPb(in)

	err = u.driverRepo.Create(ctx)
	if err != nil {
		return nil, err
	}

	return u.driverRepo.GetPb(), nil
}
