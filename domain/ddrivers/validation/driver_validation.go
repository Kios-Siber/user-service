package validation

import (
	"log"
	"skeleton/domain/ddrivers"
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
