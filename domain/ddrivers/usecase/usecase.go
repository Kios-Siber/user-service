package usecase

import (
	"log"
	"skeleton/domain/ddrivers"
)

type service struct {
	log        *log.Logger
	driverRepo ddrivers.DriverRepoInterface
}

func NewService(log *log.Logger, driverRepo ddrivers.DriverRepoInterface) ddrivers.DriverUsecaseInterface {
	return &service{
		log:        log,
		driverRepo: driverRepo,
	}
}
