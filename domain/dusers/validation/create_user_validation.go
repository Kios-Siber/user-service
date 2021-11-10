package validation

import (
	"context"
	"ksuser/lib/helper"
	"ksuser/pb/users"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *validation) Create(ctx context.Context, in *users.User) error {
	select {
	case <-ctx.Done():
		return helper.ContextError(ctx)
	default:
	}

	if len(in.Name) == 0 {
		u.log.Println("please supply valid name")
		return status.Error(codes.InvalidArgument, "please supply valid name")
	}

	if len(in.Username) == 0 {
		u.log.Println("please supply valid username")
		return status.Error(codes.InvalidArgument, "please supply valid username")
	}

	if len(in.Email) == 0 {
		u.log.Println("please supply valid email")
		return status.Error(codes.InvalidArgument, "please supply valid email")
	}

	return nil
}
