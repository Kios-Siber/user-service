package validation

import (
	"context"
	"ksuser/lib/helper"
	"ksuser/pb/employees"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *employeeValidation) Create(ctx context.Context, in *employees.Employee) error {
	select {
	case <-ctx.Done():
		return helper.ContextError(ctx)
	default:
	}

	if len(in.Name) == 0 {
		u.log.Println("please supply valid name")
		return status.Error(codes.InvalidArgument, "please supply valid name")
	}

	if len(in.Code) == 0 {
		u.log.Println("please supply valid code")
		return status.Error(codes.InvalidArgument, "please supply valid code")
	}

	if len(in.User.Id) == 0 {
		u.log.Println("please supply valid user")
		return status.Error(codes.InvalidArgument, "please supply valid user")
	}

	if len(in.Address) == 0 {
		u.log.Println("please supply valid address")
		return status.Error(codes.InvalidArgument, "please supply valid address")
	}

	if len(in.CityId) == 0 {
		u.log.Println("please supply valid city ID")
		return status.Error(codes.InvalidArgument, "please supply valid city ID")
	}

	if len(in.City) == 0 {
		u.log.Println("please supply valid city")
		return status.Error(codes.InvalidArgument, "please supply valid city")
	}

	if len(in.ProvinceId) == 0 {
		u.log.Println("please supply valid province ID")
		return status.Error(codes.InvalidArgument, "please supply valid province")
	}

	if len(in.Jabatan) == 0 {
		u.log.Println("please supply valid jabatan")
		return status.Error(codes.InvalidArgument, "please supply valid jabatan")
	}

	return nil
}
