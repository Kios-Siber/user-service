package repositories

import (
	"context"
	"skeleton/lib/helper"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *repo) Create(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return helper.ContextError(ctx)
	default:
	}

	query := `
        INSERT INTO drivers (
            id, name, phone, licence_number, company_id, company_name, created, created_by, updated, updated_by)
        VALUES ($1, $2, $3 ,$4, $5, $6, $7, $8, $9, $10)
    `
	u.pb.Id = uuid.New().String()
	now := time.Now().Format("2006-01-02 15:04:05.000000")
	_, err := u.db.ExecContext(ctx, query,
		u.pb.Id, u.pb.Name, u.pb.Phone, u.pb.LicenceNumber, u.pb.CompanyId, u.pb.CompanyName,
		now, u.pb.CreatedBy, now, u.pb.UpdatedBy)

	if err != nil {
		u.log.Println(err.Error())
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
