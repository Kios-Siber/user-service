package repositories

import (
	"database/sql"
	"log"
	"skeleton/domain/ddrivers"
	"skeleton/pb/drivers"
)

type repo struct {
	db  *sql.DB
	log *log.Logger
	pb  drivers.Driver
}

func NewDriverRepo(db *sql.DB, log *log.Logger) ddrivers.DriverRepoInterface {
	return &repo{
		db:  db,
		log: log,
	}
}

func (u *repo) GetPb() *drivers.Driver {
	return &u.pb
}

func (u *repo) SetPb(in *drivers.Driver) {
	if len(in.Id) > 0 {
		u.pb.Id = in.Id
	}
	if len(in.Name) > 0 {
		u.pb.Name = in.Name
	}
	if len(in.Phone) > 0 {
		u.pb.Phone = in.Phone
	}
	if len(in.LicenceNumber) > 0 {
		u.pb.LicenceNumber = in.LicenceNumber
	}
	if len(in.CompanyId) > 0 {
		u.pb.CompanyId = in.CompanyId
	}
	if len(in.CompanyName) > 0 {
		u.pb.CompanyName = in.CompanyName
	}
	u.pb.IsDelete = in.IsDelete
	if len(in.Created) > 0 {
		u.pb.Created = in.Created
	}
	if len(in.CreatedBy) > 0 {
		u.pb.CreatedBy = in.CreatedBy
	}
	if len(in.Updated) > 0 {
		u.pb.Updated = in.Updated
	}
	if len(in.UpdatedBy) > 0 {
		u.pb.UpdatedBy = in.UpdatedBy
	}
}
