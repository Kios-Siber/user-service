package route

import (
	"database/sql"
	driverHandler "ksuser/domain/ddrivers/handler"
	driverRepo "ksuser/domain/ddrivers/repositories"
	driverUsecase "ksuser/domain/ddrivers/usecase"
	"ksuser/pb/drivers"
	"log"

	"google.golang.org/grpc"
)

func GrpcRoute(grpcServer *grpc.Server, log *log.Logger, db *sql.DB) {
	driverServer := driverHandler.NewDriverHandler(
		driverUsecase.NewService(log, driverRepo.NewDriverRepo(db, log)),
	)

	drivers.RegisterDriversServiceServer(grpcServer, driverServer)
}
