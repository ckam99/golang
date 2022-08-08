package main

import (
	"log"

	"github.com/ckam225/golang/echo/internal/database/postgres/storage"
	"github.com/ckam225/golang/echo/internal/handler"
)

func main() {
	store, err := storage.NewStore("postgres://postgres:postgres@host.docker.internal/golang_echo?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	server := handler.NewHandler(*store)
	server.Logger.Fatal(server.Start(":8000"))
}
