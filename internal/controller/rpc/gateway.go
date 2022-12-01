package rpc

import (
	"context"
	"example/grpc/internal/controller/rpc/handler"
	"example/grpc/internal/controller/rpc/pb"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// Serve starts gRPC server
func (s *Server) ServeHttpGateway(address string) error {
	rmux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// register all grpc service handlers
	if err := pb.RegisterAuthorServiceHandlerServer(ctx, rmux,
		handler.NewAuthorServer(s.db, s.logger)); err != nil {
		return fmt.Errorf("cannot register author handler server: %s", err)
	}
	if err := pb.RegisterBookServiceHandlerServer(ctx, rmux,
		handler.NewBookServer(s.db, s.logger)); err != nil {
		return fmt.Errorf("cannot register book handler server: %s", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", rmux)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("cannot create http gateway network listener:%w", err)
	}
	if err = http.Serve(listener, mux); err != nil {
		return fmt.Errorf("cannot start HTTP gateway server: %w", err)
	}
	return nil
}
