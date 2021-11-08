package repositories

import (
	"context"
	"skeleton/lib/helper"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *repo) Update(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return helper.ContextError(ctx)
	default:
	}

	query := `
        UPDATE drivers 
        SET name = $1, 
                phone = $2, 
                licence_number = $3, 
                updated = $4, 
                updated_by = $5
        WHERE id = $6
    `
	now := time.Now().Format("2006-01-02 15:04:05.000000")
	_, err := u.db.ExecContext(ctx, query,
		u.pb.Name, u.pb.Phone, u.pb.LicenceNumber, now, u.pb.UpdatedBy, u.pb.Id)

	if err != nil {
		u.log.Println(err.Error())
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
