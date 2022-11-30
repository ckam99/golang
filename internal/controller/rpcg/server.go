package rpcg

import (
	"example/grpc/internal/controller/rpcg/protobuf"
	"example/grpc/internal/core/ports"
	"example/grpc/internal/core/service"
	"example/grpc/internal/provider/postgres"
	"example/grpc/pkg/postgresql"
)

// Server serves gRPC requests for core business logics services.
type Server struct {
	protobuf.UnimplementedBookServiceServer
	bookService ports.BookService
}

// NewServer created a new gRPC server.
func NewServer(db postgresql.Client) *Server {
	return &Server{
		bookService: service.NewBookService(
			postgres.NewBookRepository(db),
		),
	}
}
