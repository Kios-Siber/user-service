package handler

import (
	"context"
	"ksuser/domain/demployees"
	"ksuser/pb/employees"
)

type EmployeeHandler struct {
	usecase demployees.EmployeeUsecaseInterface
}

func NewEmployeeHandler(usecase demployees.EmployeeUsecaseInterface) *EmployeeHandler {
	handler := new(EmployeeHandler)
	handler.usecase = usecase
	return handler
}

func (u *EmployeeHandler) Create(ctx context.Context, in *employees.Employee) (*employees.Employee, error) {
	return u.usecase.Create(ctx, in)
}
