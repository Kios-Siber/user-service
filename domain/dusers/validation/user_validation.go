package validation

import (
	"ksuser/domain/dusers"
	"log"
)

type validation struct {
	log      *log.Logger
	userRepo dusers.UserRepoInterface
}

func NewUserValidation(log *log.Logger, userRepo dusers.UserRepoInterface) dusers.UserValidationInterface {
	return &validation{
		log:      log,
		userRepo: userRepo,
	}
}
