package handler

import (
	"context"
	"skeleton/domain/ddrivers"
	"skeleton/pb/drivers"
	"skeleton/pb/generic"
)

type DriverHandler struct {
	usecase ddrivers.DriverUsecaseInterface
}

func NewDriverHandler(usecase ddrivers.DriverUsecaseInterface) *DriverHandler {
	handler := new(DriverHandler)
	handler.usecase = usecase
	return handler
}

func (u *DriverHandler) List(ctx context.Context, in *drivers.DriverListInput) (*drivers.Drivers, error) {
	return u.usecase.List(ctx, in)
}

func (u *DriverHandler) Create(ctx context.Context, in *drivers.Driver) (*drivers.Driver, error) {
	return u.usecase.Create(ctx, in)
}

func (u *DriverHandler) Update(ctx context.Context, in *drivers.Driver) (*drivers.Driver, error) {
	return u.usecase.Update(ctx, in)
}

func (u *DriverHandler) Delete(ctx context.Context, in *generic.Id) (*generic.BoolMessage, error) {
	return u.usecase.Delete(ctx, in)
}
