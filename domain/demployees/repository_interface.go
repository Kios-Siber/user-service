package demployees

import (
	"context"
	"ksuser/pb/employees"
)

type EmployeeRepoInterface interface {
	Create(ctx context.Context) error
	GetPb() *employees.Employee
	SetPb(*employees.Employee)
}
