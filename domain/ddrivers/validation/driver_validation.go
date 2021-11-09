package validation

import (
	"ksuser/domain/ddrivers"
	"log"
)

type driverValidation struct {
	log        *log.Logger
	driverRepo ddrivers.DriverRepoInterface
}

func NewValidation(log *log.Logger, driverRepo ddrivers.DriverRepoInterface) ddrivers.DriverValidationInterface {
	return &driverValidation{
		log:        log,
		driverRepo: driverRepo,
	}
}
