package rpc

import (
	"example/grpc/internal/controller/rpc/handler"
	"example/grpc/internal/controller/rpc/pb"
	"example/grpc/pkg/postgresql"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server serves gRPC requests for core business logics services.
type Server struct {
	logger *log.Logger
	db     postgresql.Client
}

// NewServer created a new gRPC server.
func NewServer(db postgresql.Client, logger *log.Logger) *Server {
	return &Server{
		logger: logger,
		db:     db,
	}
}

// Serve starts gRPC server
func (s *Server) Serve(address string) error {
	gRPCServer := grpc.NewServer()
	// register all grpc service here
	pb.RegisterAuthorServiceServer(gRPCServer, handler.NewAuthorServer(s.db, s.logger))
	pb.RegisterBookServiceServer(gRPCServer, handler.NewBookServer(s.db, s.logger))
	reflection.Register(gRPCServer)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("cannot create network listener:%w", err)
	}
	if err = gRPCServer.Serve(listener); err != nil {
		return fmt.Errorf("cannot start gRPC server: %w", err)
	}
	return nil
}
