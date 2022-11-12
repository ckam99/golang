package main

import (
	"app/internal/domain/author"
	"app/pkg/clients/postgresql"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	pg, err := postgresql.NewClient(context.TODO(), postgresql.Config{
		Host:     "host.docker.internal",
		Port:     "54323",
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
