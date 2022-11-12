package main

import (
	"log"
	"main/internal/config"
	"main/internal/controller/http"
	"main/pkg/clients/sqlite"
)

func main() {
	cfg := config.Config{}
	db, err := sqlite.Connect(":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	server := http.NewHTTP(db, cfg)
	log.Fatal(server.Run())
}
