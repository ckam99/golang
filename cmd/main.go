package main

import (
	"log"
	"main/internal/config"
	"main/internal/controller/http"
	"main/pkg/clients/sqlite"
	"main/pkg/migrate"
)

func main() {
	cfg := config.Config{}
	db, err := sqlite.Connect("example.db")
	if err != nil {
		panic(err)
	}
	m, err := migrate.New("./internal/migration", "sqlite3", "example.db", &migrate.Config{})
	if err != nil {
		panic(err)
	}
	if err = m.Migrate(); err != nil {
		panic(err)
	}
	defer db.Close()
	server := http.NewHTTP(db, cfg)
	log.Fatal(server.Run())
}
