package main

import (
	"app/internal/domain/author"
	"app/pkg/clients/postgresql"
	"app/pkg/migrate"
	"context"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	m, err := migrate.New("./migrations", "postgres",
		"postgres://postgres:postgres@host.docker.internal/demo?sslmode=disable",
		&migrate.Config{},
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Migrate(); err != nil {
		panic(err)
	}
}

func mainw() {

	pg, err := postgresql.NewClient(context.TODO(), postgresql.Config{
		Host:     "host.docker.internal",
		Port:     "5432",
		Database: "demo",
		Password: "postgres",
		Username: "postgres",
	}, 3)
	if err != nil {
		log.Fatal(err)
	}
	repository := author.NewRepository(pg)

	authors, err := repository.GetAll(context.TODO(), author.FilterParamsDTO{
		Limit:  5,
		Offset: 5,
		// OrderBy:   []string{"id", "name"},
		// Ascending: "desc",
	})

	if err != nil {
		log.Fatal(err)
	}
	PrintJSON(authors)
}

func PrintJSON(d interface{}) {
	b, err := json.MarshalIndent(&d, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
