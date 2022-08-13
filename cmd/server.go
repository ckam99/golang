package main

import (
	"fmt"
	"log"

	"github.com/ckam225/golang/fiber-sqlx/internal/config"
	"github.com/ckam225/golang/fiber-sqlx/internal/database/postgres/storage"
	"github.com/ckam225/golang/fiber-sqlx/internal/handler"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	//store, err := storage.NewStore("postgres://postgres:postgres@host.docker.internal:5432/golang_echo?sslmode=disable")
	store, err := storage.NewStore(cfg.Database.GetURL())

	if err != nil {
		log.Fatal(err)
	}
	server := handler.NewHandler(*store)
	log.Fatal(server.Listen(fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)))
}
