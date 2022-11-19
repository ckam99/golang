package postgres

import (
	"context"
	"log"
	"main/pkg/client/postgresql"
)

type Client struct {
	postgresql.Client
}

func New(dsn string) Client {
	c, err := postgresql.NewClient(context.Background(), dsn, 3)
	if err != nil {
		log.Fatal(err)
	}
	return Client{
		Client: c,
	}
}
