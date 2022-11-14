package main

import (
	"context"
	"main/internal/adapter/database"
	"main/pkg/clients/postgresql"
)

func main() {
	//cfg := config.Config{}

	db, err := database.Connection(context.Background(), postgresql.Config{
		Host:     "host.docker.internal",
		Username: "postgres",
		Password: "postgres",
		Database: "demo",
		Port:     "5432",
		SSLMode:  "disable",
	})
	if err != nil {
		panic(err)
	}
	if err = db.Rollback(); err != nil {
		panic(err)
	}

	defer db.Close()

	//server := http.NewHTTP(db, cfg)
	//log.Fatal(server.Run())
}
