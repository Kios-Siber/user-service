package usecase

import (
	"context"
	"skeleton/lib/helper"
	"skeleton/pb/drivers"
)

func (u *service) List(ctx context.Context, in *drivers.DriverListInput) (*drivers.Drivers, error) {
	select {
	case <-ctx.Done():
		return nil, helper.ContextError(ctx)
	default:
	}

	return u.driverRepo.FindAll(ctx, in)
}
