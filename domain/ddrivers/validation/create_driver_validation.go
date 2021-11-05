package validation

import (
	"context"
	"skeleton/lib/helper"
	"skeleton/pb/drivers"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *driverValidation) Create(ctx context.Context, in *drivers.Driver) error {
	select {
	case <-ctx.Done():
		return helper.ContextError(ctx)
	default:
	}

	if len(in.Name) == 0 {
		u.log.Println("please supply valid name")
		return status.Error(codes.InvalidArgument, "please supply valid name")
	}

	if len(in.Phone) == 0 {
		u.log.Println("please supply valid phone")
		return status.Error(codes.InvalidArgument, "please supply valid phone")
	}

	if len(in.CompanyId) == 0 {
		u.log.Println("please supply valid company id")
		return status.Error(codes.InvalidArgument, "please supply valid company id")
	}

	if len(in.CompanyName) == 0 {
		u.log.Println("please supply valid company name")
		return status.Error(codes.InvalidArgument, "please supply valid company name")
	}

	if len(in.LicenceNumber) == 0 {
		u.log.Println("please supply valid licence number")
		return status.Error(codes.InvalidArgument, "please supply valid licence number")
	}

	return nil
}
