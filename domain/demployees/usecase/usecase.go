package usecase

import (
	"ksuser/domain/demployees"
	"log"
)

type service struct {
	log          *log.Logger
	employeeRepo demployees.EmployeeRepoInterface
}

func NewService(log *log.Logger, employeeRepo demployees.EmployeeRepoInterface) demployees.EmployeeUsecaseInterface {
	return &service{
		log:          log,
		employeeRepo: employeeRepo,
	}
}
