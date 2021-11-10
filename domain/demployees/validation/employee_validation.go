package validation

import (
	"ksuser/domain/demployees"
	"log"
)

type employeeValidation struct {
	log          *log.Logger
	employeeRepo demployees.EmployeeRepoInterface
}

func NewValidation(log *log.Logger, employeeRepo demployees.EmployeeRepoInterface) demployees.EmployeeValidationInterface {
	return &employeeValidation{
		log:          log,
		employeeRepo: employeeRepo,
	}
}
