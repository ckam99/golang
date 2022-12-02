package rpc

import (
	"example/grpc/internal/controller/rpc/handler"
	"example/grpc/internal/controller/rpc/interceptor"
	"example/grpc/internal/controller/rpc/pb"
	"example/grpc/pkg/postgresql"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server serves gRPC requests for core business logics services.
type Server struct {
	db postgresql.Client
}

// NewServer created a new gRPC server.
func NewServer(db postgresql.Client) *Server {
	return &Server{
		db: db,
	}
}

// Serve starts gRPC server
func (s *Server) Serve(address string) error {

	grpcLogger := grpc.UnaryInterceptor(interceptor.GrpcLogger)
	gRPCServer := grpc.NewServer(grpcLogger)
	// register all grpc service here
	pb.RegisterAuthorServiceServer(gRPCServer, handler.NewAuthorServer(s.db))
	pb.RegisterBookServiceServer(gRPCServer, handler.NewBookServer(s.db))
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
