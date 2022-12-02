package rpc

import (
	"context"
	"example/grpc/internal/controller/http"
	"example/grpc/internal/controller/rpc/handler"
	"example/grpc/internal/controller/rpc/pb"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/encoding/protojson"
)

// Serve starts gRPC server
func (s *Server) ServeHttpGateway(address string) error {
	muxJsonOptions := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames:   true,
			EmitUnpopulated: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})
	rmux := runtime.NewServeMux(muxJsonOptions)
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

	mux := http.NewHTTPServer()
	mux.Handle("/", rmux)
	if err := mux.Serve(address); err != nil {
		return fmt.Errorf("grpc gateway :%w", err)
	}

	// mux := http.NewServeMux()
	// mux.Handle("/", rmux)

	// mux := http.NewHTTP

	// listener, err := net.Listen("tcp", address)
	// if err != nil {
	// 	return fmt.Errorf("cannot create http gateway network listener:%w", err)
	// }
	// if err = http.Serve(listener, mux); err != nil {
	// 	return fmt.Errorf("cannot start HTTP gateway server: %w", err)
	// }
	return nil
}
