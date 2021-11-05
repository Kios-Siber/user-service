package usecase

import (
	"context"
	"skeleton/domain/ddrivers/validation"
	"skeleton/lib/helper"
	"skeleton/pb/drivers"
)

func (u *service) Update(ctx context.Context, in *drivers.Driver) (*drivers.Driver, error) {
	select {
	case <-ctx.Done():
		return nil, helper.ContextError(ctx)
	default:
	}

	dValidation := validation.NewValidation(u.log, u.driverRepo)
	err := dValidation.Update(ctx, in)
	if err != nil {
		return nil, err
	}

	u.driverRepo.SetPb(in)

	err = u.driverRepo.Update(ctx)
	if err != nil {
		return nil, err
	}

	return u.driverRepo.GetPb(), nil
}
