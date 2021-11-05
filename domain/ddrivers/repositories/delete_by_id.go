package repositories

import (
	"context"
	"skeleton/lib/helper"
	"skeleton/pb/generic"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *repo) Delete(ctx context.Context, in *generic.Id) error {
	select {
	case <-ctx.Done():
		return helper.ContextError(ctx)
	default:
	}

	query := `
        UPDATE drivers 
        SET is_deleted = true
        WHERE id = $1
    `
	_, err := u.db.ExecContext(ctx, query, in.Id)

	if err != nil {
		u.log.Println(err.Error())
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
