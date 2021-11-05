package repositories

import (
	"context"
	"skeleton/lib/helper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *repo) Find(ctx context.Context, id string) error {
	select {
	case <-ctx.Done():
		return helper.ContextError(ctx)
	default:
	}

	query := `
        SELECT id, name, phone, licence_number, company_id, company_name 
        FROM drivers WHERE id = $1 AND is_deleted = false
    `

	err := u.db.QueryRowContext(ctx, query, id).Scan(
		&u.pb.Id, &u.pb.Name, &u.pb.LicenceNumber, &u.pb.CompanyId, &u.pb.CompanyName)

	if err != nil {
		u.log.Println(err.Error())
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
