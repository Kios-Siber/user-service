package demployees

import (
	"context"
	"ksuser/pb/employees"
)

type EmployeeValidationInterface interface {
	Create(ctx context.Context, id *employees.Employee) error
}
