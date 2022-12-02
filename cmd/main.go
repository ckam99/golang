package main

import (
	"context"
	"example/grpc/internal/controller/rpc"
	"example/grpc/pkg/postgresql"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	if os.Getenv("ENVIRONMENT") != "production" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	dsn := "postgres://postgres:postgres@host.docker.internal/demo?sslmode=disable"
	db, err := postgresql.Connection(ctx, dsn, 3)
	if err != nil {
		log.Info().Msgf("cannot connect to the database: %w", err)
	}
	go runHttpGatewayServer(db)
	runGrpcServer(db)
}

func runGrpcServer(db postgresql.Client) {
	server := rpc.NewServer(db)
	port := ":5000"
	log.Info().Msgf("gRPC server started at 0.0.0.0%s\n", port)
	if err := server.Serve(port); err != nil {
		log.Error().Err(err)
	}
}

func runHttpGatewayServer(db postgresql.Client) {
	server := rpc.NewServer(db)
	port := ":8000"
	log.Info().Msgf("HTTP gateway server started at 0.0.0.0%s\n", port)
	if err := server.ServeHttpGateway(port); err != nil {
		log.Error().Err(err)
	}
}
