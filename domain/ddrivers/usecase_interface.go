package ddrivers

import (
	"context"
	"ksuser/pb/drivers"
	"ksuser/pb/generic"
)

type DriverUsecaseInterface interface {
	List(ctx context.Context, in *drivers.DriverListInput) (*drivers.Drivers, error)
	Create(ctx context.Context, in *drivers.Driver) (*drivers.Driver, error)
	Update(ctx context.Context, in *drivers.Driver) (*drivers.Driver, error)
	Delete(ctx context.Context, in *generic.Id) (*generic.BoolMessage, error)
}
