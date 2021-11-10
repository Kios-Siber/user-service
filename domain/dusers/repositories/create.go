package repositories

import (
	"context"
	"ksuser/lib/app"
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
        INSERT INTO users (
					password, company_id, region_id, branch_id, username, name, email, created_by, updated_by)
        VALUES ('', $1, $2, $3 ,$4, $5, $6, $7, $8)
    `

	var branchId, regionId *string

	if u.pb.Branch != nil {
		branchId = &u.pb.Branch.Id
	}

	if u.pb.Branch != nil {
		branchId = &u.pb.Branch.Id
	}

	println(query)
	println(ctx.Value(app.Ctx("companyID")).(string), regionId, branchId, u.pb.Username, u.pb.Name, u.pb.Email, ctx.Value(app.Ctx("userID")).(string), ctx.Value(app.Ctx("userID")).(string))
	_, err := u.db.ExecContext(ctx, query,
		ctx.Value(app.Ctx("companyID")).(string), regionId, branchId, u.pb.Username, u.pb.Name, u.pb.Email, ctx.Value(app.Ctx("userID")).(string), ctx.Value(app.Ctx("userID")).(string))

	if err != nil {
		u.log.Println(err.Error())
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
