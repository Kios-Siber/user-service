package main

import (
	"log"
	"net"
	"os"
	"skeleton/config"
	"skeleton/lib/database/postgres"
	"skeleton/route"

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
	route.GrpcRoute(grpcServer, log, db)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
		return
	}
	log.Print("serve grpc on port: " + os.Getenv("PORT"))

}
