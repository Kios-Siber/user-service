package ddrivers

import (
	"context"
	"ksuser/pb/drivers"
	"ksuser/pb/generic"
)

type DriverRepoInterface interface {
	Find(ctx context.Context, id string) error
	FindAll(ctx context.Context, in *drivers.DriverListInput) (*drivers.Drivers, error)
	Create(ctx context.Context) error
	Update(ctx context.Context) error
	Delete(ctx context.Context, in *generic.Id) error
	GetPb() *drivers.Driver
	SetPb(*drivers.Driver)
}
