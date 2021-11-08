package route

import (
	"database/sql"
	"log"
	driverHandler "skeleton/domain/ddrivers/handler"
	driverRepo "skeleton/domain/ddrivers/repositories"
	driverUsecase "skeleton/domain/ddrivers/usecase"
	"skeleton/pb/drivers"

	"google.golang.org/grpc"
)

func GrpcRoute(grpcServer *grpc.Server, log *log.Logger, db *sql.DB) {
	driverServer := driverHandler.NewDriverHandler(
		driverUsecase.NewService(log, driverRepo.NewDriverRepo(db, log)),
	)

	drivers.RegisterDriversServiceServer(grpcServer, driverServer)
}
