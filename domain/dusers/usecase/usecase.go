package usecase

import (
	"ksuser/domain/dusers"
	"log"
)

type userservice struct {
	log      *log.Logger
	userRepo dusers.UserRepoInterface
}

func NewUserService(log *log.Logger, userRepo dusers.UserRepoInterface) dusers.UserUsecaseInterface {
	return &userservice{
		log:      log,
		userRepo: userRepo,
	}
}
