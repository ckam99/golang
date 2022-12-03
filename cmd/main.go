package main

import (
	"example/grpc/controller"
	pb "example/grpc/pb/service"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	gRPCServer := grpc.NewServer()
	authorServer := controller.NewAuthorServer()
	// register all grpc service here
	pb.RegisterAuthorServiceServer(gRPCServer, authorServer)
	reflection.Register(gRPCServer)
	listener, err := net.Listen("tcp", ":6000")
	if err != nil {
		log.Fatalf("cannot create network listener:%s", err)
	}
	log.Println("gRPC server start on 0.0.0.0:6000")
	if err = gRPCServer.Serve(listener); err != nil {
		log.Fatalf("cannot start gRPC server: %s", err)
	}
}
