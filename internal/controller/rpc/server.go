package rpc

import (
	"example/grpc/internal/controller/rpc/protobuf"
	"example/grpc/internal/core/ports"
	"example/grpc/internal/core/service"
	"example/grpc/internal/provider/postgres"
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
	protobuf.UnimplementedAppServiceServer
	authorService ports.AuthorService
	bookService   ports.BookService
}

// NewServer created a new gRPC server.
func NewServer(db postgresql.Client, logger *log.Logger) *Server {
	return &Server{
		logger: logger,
		authorService: service.NewAuthorService(
			postgres.NewAuthorRepository(db),
		),
		bookService: service.NewBookService(
			postgres.NewBookRepository(db),
		),
	}
}

// Serve starts gRPC server
func (s *Server) Serve(host string) error {
	gRPCServer := grpc.NewServer()
	protobuf.RegisterAppServiceServer(gRPCServer, s)
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
