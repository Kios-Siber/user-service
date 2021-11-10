package repositories

import (
	"database/sql"
	"ksuser/domain/dusers"
	"ksuser/pb/users"
	"log"
)

type repo struct {
	db  *sql.DB
	log *log.Logger
	pb  users.User
}

func NewUserRepo(db *sql.DB, log *log.Logger) dusers.UserRepoInterface {
	return &repo{
		db:  db,
		log: log,
	}
}

func (u *repo) GetPb() *users.User {
	return &u.pb
}

func (u *repo) SetPb(in *users.User) {
	if len(in.Id) > 0 {
		u.pb.Id = in.Id
	}
	if len(in.Name) > 0 {
		u.pb.Name = in.Name
	}
	if in.Region != nil {
		u.pb.Region = in.Region
	}
	if in.Branch != nil {
		u.pb.Branch = in.Branch
	}
	if len(in.Username) > 0 {
		u.pb.Username = in.Username
	}
	if len(in.Email) > 0 {
		u.pb.Email = in.Email
	}
	if len(in.UpdatedBy) > 0 {
		u.pb.UpdatedBy = in.UpdatedBy
	}
	if len(in.CreatedAt) > 0 {
		u.pb.CreatedAt = in.CreatedAt
	}
	if len(in.UpdatedAt) > 0 {
		u.pb.UpdatedAt = in.UpdatedAt
	}
}
