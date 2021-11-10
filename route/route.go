package route

import (
	"database/sql"
	"ksuser/pb/users"
	"log"

	userHandler "ksuser/domain/dusers/handler"
	userRepo "ksuser/domain/dusers/repositories"
	userUsecase "ksuser/domain/dusers/usecase"

	"google.golang.org/grpc"
)

func GrpcRoute(grpcServer *grpc.Server, log *log.Logger, db *sql.DB) {
	userServer := userHandler.NewUserHandler(
		userUsecase.NewUserService(log, userRepo.NewUserRepo(db, log)),
	)

	users.RegisterUsersServiceServer(grpcServer, userServer)
}
