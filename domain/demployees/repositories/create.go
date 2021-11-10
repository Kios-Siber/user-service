package repositories

import (
	"context"
	"ksuser/lib/helper"

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
        INSERT INTO employees (
            name, user_id, code, address, city_id, city, province_id, province, jabatan, created_by, updated_by)
        VALUES ($1, $2, $3 ,$4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
    `
	_, err := u.db.ExecContext(ctx, query,
		u.pb.Name, u.pb.User.Id, u.pb.Code, u.pb.Address, u.pb.CityId, u.pb.City, &u.pb.ProvinceId, &u.pb.Province,
		u.pb.Jabatan, u.pb.CreatedBy, u.pb.UpdatedBy)

	if err != nil {
		u.log.Println(err.Error())
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
