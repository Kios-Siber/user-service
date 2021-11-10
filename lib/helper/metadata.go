package helper

import (
	"context"
	"ksuser/lib/app"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func GetMetadata(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	userID := md["user_id"]
	if len(userID) == 0 {
		return ctx, status.Errorf(codes.Unauthenticated, "user_id is not provided")
	}

	ctx = context.WithValue(ctx, app.Ctx("userID"), userID[0])

	companyID := md["company_id"]
	if len(companyID) == 0 {
		return ctx, status.Errorf(codes.Unauthenticated, "company_id is not provided")
	}

	ctx = context.WithValue(ctx, app.Ctx("companyID"), companyID[0])

	branch := md["branch_id"]
	var branchID string
	if branch != nil {
		branchID = branch[0]
	}
	ctx = context.WithValue(ctx, app.Ctx("branchID"), branchID)

	region := md["region_id"]
	var regionID string
	if region != nil {
		regionID = region[0]
	}
	ctx = context.WithValue(ctx, app.Ctx("regionID"), regionID)

	return ctx, nil
}
