package usecase

import (
	"context"
	"ksuser/domain/demployees/validation"
	"ksuser/lib/helper"
	"ksuser/pb/employees"
)

func (u *service) Create(ctx context.Context, in *employees.Employee) (*employees.Employee, error) {
	select {
	case <-ctx.Done():
		return nil, helper.ContextError(ctx)
	default:
	}

	dValidation := validation.NewValidation(u.log, u.employeeRepo)
	err := dValidation.Create(ctx, in)
	if err != nil {
		return nil, err
	}

	// in.CreatedBy = metadata.user_id
	// in.UpdatedBy = metadata.user_id

	u.employeeRepo.SetPb(in)

	err = u.employeeRepo.Create(ctx)
	if err != nil {
		return nil, err
	}

	return u.employeeRepo.GetPb(), nil
}
