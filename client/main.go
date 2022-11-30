package main

import (
	"context"
	"example/grpc/internal/controller/rpc/protobuf"
	"example/grpc/pkg/postgresql"
	"example/grpc/pkg/utils"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	dsn := "postgres://postgres:postgres@host.docker.internal/demo?sslmode=disable"
	db, err := postgresql.Connection(ctx, dsn, 3)
	if err != nil {
		log.Fatal("cannot connect to the database: ", err)
	}
	runGrpcClient(db)
}

func runGrpcClient(db postgresql.Client) {
	conn, err := grpc.Dial(":8000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := protobuf.NewBookServiceClient(conn)
	runfetchBooksStreamer(client)
}

func runfetchBooksStreamer(client protobuf.BookServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := client.GetBooks(ctx, &protobuf.QueryRequest{})
	if err != nil {
		log.Fatal(err)
	}
	for {
		row, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Panic(err)
		}
		utils.JSON(row)
	}
}
