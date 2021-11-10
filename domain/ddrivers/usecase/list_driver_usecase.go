package usecase

import (
	"context"
	"ksuser/lib/helper"
	"ksuser/pb/drivers"
)

func (u *service) List(ctx context.Context, in *drivers.DriverListInput) (*drivers.Drivers, error) {
	select {
	case <-ctx.Done():
		return nil, helper.ContextError(ctx)
	default:
	}

	return u.driverRepo.FindAll(ctx, in)
}
