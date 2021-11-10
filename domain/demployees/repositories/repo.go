package repositories

import (
	"database/sql"
	"ksuser/domain/demployees"
	"ksuser/pb/employees"
	"log"
)

type repo struct {
	db  *sql.DB
	log *log.Logger
	pb  employees.Employee
}

func NewEmployeeRepo(db *sql.DB, log *log.Logger) demployees.EmployeeRepoInterface {
	return &repo{
		db:  db,
		log: log,
	}
}

func (u *repo) GetPb() *employees.Employee {
	return &u.pb
}

func (u *repo) SetPb(in *employees.Employee) {
	if len(in.Id) > 0 {
		u.pb.Id = in.Id
	}
	if len(in.Name) > 0 {
		u.pb.Name = in.Name
	}
	if in.User != nil && len(in.User.Id) > 0 {
		u.pb.User = in.User
	}
	if len(in.Code) > 0 {
		u.pb.Code = in.Code
	}
	if len(in.Address) > 0 {
		u.pb.Address = in.Address
	}
	if len(in.CityId) > 0 {
		u.pb.CityId = in.CityId
	}
	if len(in.City) > 0 {
		u.pb.City = in.City
	}
	if len(in.ProvinceId) > 0 {
		u.pb.ProvinceId = in.ProvinceId
	}
	if len(in.Province) > 0 {
		u.pb.Province = in.Province
	}
	if len(in.Jabatan) > 0 {
		u.pb.Jabatan = in.Jabatan
	}
	if len(in.CreatedBy) > 0 {
		u.pb.CreatedBy = in.CreatedBy
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
