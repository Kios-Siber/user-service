package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"os"
	"skeleton/config"
	"skeleton/lib/database/postgres"
	"skeleton/pb/drivers"
	"skeleton/pb/generic"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	config.Setup(".env")

	log := log.New(os.Stdout, "grpc skeleton : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	db, err := postgres.Open()
	if err != nil {
		log.Fatalf("connecting to db: %v", err)
		return
	}
	log.Print("connecting to postgresql database")
	defer db.Close()

	// listen tcp port
	lis, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	grpcServer := grpc.NewServer()

	// routing grpc services
	grpcRoute(grpcServer, log, db)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
		return
	}
	log.Print("serve grpc on port: " + os.Getenv("PORT"))

}

func grpcRoute(grpcServer *grpc.Server, log *log.Logger, db *sql.DB) {
	driverServer := newDriverHandler(log, db)

	drivers.RegisterDriversServiceServer(grpcServer, driverServer)
}

type driverHandler struct {
	log *log.Logger
	db  *sql.DB
}

func newDriverHandler(log *log.Logger, db *sql.DB) *driverHandler {
	handler := new(driverHandler)
	handler.log = log
	handler.db = db
	return handler
}

func (u *driverHandler) List(ctx context.Context, in *drivers.DriverListInput) (*drivers.Drivers, error) {
	return &drivers.Drivers{}, nil
}

func (u *driverHandler) Create(ctx context.Context, in *drivers.Driver) (*drivers.Driver, error) {
	query := `
        INSERT INTO drivers (
            id, name, phone, licence_number, company_id, company_name, created, created_by, updated, updated_by)
        VALUES ($1, $2, $3 ,$4, $5, $6, $7, $8, $9, $10)
    `
	in.Id = uuid.New().String()
	now := time.Now().Format("2006-01-02 15:04:05.000000")
	_, err := u.db.ExecContext(ctx, query,
		in.Id, in.Name, in.Phone, in.LicenceNumber, in.CompanyId, in.CompanyName, now, "jaka", now, "jaka")
	if err != nil {
		u.log.Println(err.Error())
		return &drivers.Driver{}, err
	}
	return in, nil
}

func (u *driverHandler) Update(ctx context.Context, in *drivers.Driver) (*drivers.Driver, error) {
	return in, nil
}

func (u *driverHandler) Delete(ctx context.Context, in *generic.Id) (*generic.BoolMessage, error) {
	return &generic.BoolMessage{}, nil
}
