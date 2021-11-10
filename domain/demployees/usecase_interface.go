package demployees

import (
	"context"
	"ksuser/pb/employees"
)

type EmployeeUsecaseInterface interface {
	Create(ctx context.Context, in *employees.Employee) (*employees.Employee, error)
}
