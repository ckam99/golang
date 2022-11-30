package rpc

import (
	"example/grpc/internal/controller/rpc/handler"
	"example/grpc/internal/controller/rpc/protobuf"
	"example/grpc/pkg/postgresql"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server serves gRPC requests for core business logics services.
type Server struct {
	db     postgresql.Client
	logger *log.Logger
	*handler.AuthorServer
	*handler.BookServer
}

// NewServer created a new gRPC server.
func NewServer(db postgresql.Client, logger *log.Logger) *Server {
	return &Server{
		db:     db,
		logger: logger,
	}
}

// Serve starts gRPC server
func (s *Server) Serve(host string) error {
	gRPCServer := grpc.NewServer()
	// register all grpc service here
	protobuf.RegisterAuthorServiceServer(gRPCServer, handler.NewAuthorServer(s.db))
	protobuf.RegisterBookServiceServer(gRPCServer, handler.NewBookServer(s.db))

	reflection.Register(gRPCServer)
	if host == "" {
		host = ":8000"
	}
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return fmt.Errorf("cannot create network listener:%w", err)
	}
	if err = gRPCServer.Serve(listener); err != nil {
		return fmt.Errorf("cannot start gRPC server: %w", err)
	}
	return nil
}
