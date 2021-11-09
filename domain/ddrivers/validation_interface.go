package ddrivers

import (
	"context"
	"ksuser/pb/drivers"
)

type DriverValidationInterface interface {
	Create(ctx context.Context, id *drivers.Driver) error
	Update(ctx context.Context, id *drivers.Driver) error
	Delete(ctx context.Context, id string) error
}
