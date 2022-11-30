package main

import (
	"context"
	"example/grpc/internal/controller/rpc"
	"example/grpc/internal/core/entity"
	"example/grpc/internal/core/service"
	"example/grpc/internal/provider/postgres"
	"example/grpc/pkg/postgresql"
	"example/grpc/pkg/utils"
	"log"
	"time"
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

func runGrpcServer(db postgresql.Client) {
	server := rpc.NewServer(db, log.Default())
	log.Println("gRPC server started at 0.0.0.0:8000")
	if err := server.Serve(":8000"); err != nil {
		log.Fatal(err)
	}
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
