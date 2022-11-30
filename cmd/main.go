package main

import (
	"context"
	"encoding/json"
	"example/grpc/internal/controller/rpcg"
	"example/grpc/internal/controller/rpcg/protobuf"
	"example/grpc/internal/core/entity"
	"example/grpc/internal/core/service"
	"example/grpc/internal/provider/postgres"
	"example/grpc/pkg/postgresql"
	"example/grpc/pkg/utils"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	dsn := "postgres://postgres:postgres@host.docker.internal/demo?sslmode=disable"
	db, err := postgresql.Connection(ctx, dsn, 3)
	if err != nil {
		log.Fatal("cannot connect to the database: ", err)
	}

	// testBookService(db)

	runGrpcServer(db)
}

func testBookService(db postgresql.Client) {
	book := entity.Book{
		Title:       "jk jg678",
		Description: "kj hjjg dfdfdfsdfsdf",
		PublishedAt: &time.Time{},
	}
	bs := service.NewBookService(postgres.NewBookRepository(db))
	err := bs.Create(context.TODO(), &book)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("book created")
	utils.JSON(book)

	books, err := bs.GetAll(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	utils.JSON(books)
}

func output(t any) {
	b, _ := json.MarshalIndent(t, "", " ")
	log.Println(string(b))
}

func runGrpcServer(db postgresql.Client) {
	gRPCServer := grpc.NewServer()
	server := rpcg.NewServer(db)
	protobuf.RegisterBookServiceServer(gRPCServer, server)
	reflection.Register(gRPCServer)

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("cannot create network listener: ", err)
	}
	log.Println("gRPC server started at 0.0.0.0:8000")
	if err = gRPCServer.Serve(listener); err != nil {
		log.Fatal("cannot start gRPC server: ", err)
	}
}
