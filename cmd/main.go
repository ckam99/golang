package main

import (
	"context"
	"example/grpc/internal/controller/rpc"
	"example/grpc/pkg/postgresql"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	dsn := "postgres://postgres:postgres@host.docker.internal/demo?sslmode=disable"
	db, err := postgresql.Connection(ctx, dsn, 3)
	if err != nil {
		log.Fatal("cannot connect to the database: ", err)
	}
	go runHttpGatewayServer(db)
	runGrpcServer(db)
}

func runGrpcServer(db postgresql.Client) {
	server := rpc.NewServer(db, log.Default())
	port := ":5000"
	log.Printf("gRPC server started at 0.0.0.0%s\n", port)
	if err := server.Serve(port); err != nil {
		log.Fatal(err)
	}
}

func runHttpGatewayServer(db postgresql.Client) {
	server := rpc.NewServer(db, log.Default())
	port := ":8000"
	log.Printf("HTTP gateway server started at 0.0.0.0%s\n", port)
	if err := server.ServeHttpGateway(port); err != nil {
		log.Fatal(err)
	}
}
