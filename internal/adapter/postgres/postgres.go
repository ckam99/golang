package postgres

import (
  "main/pkg/client/postgresql"
  "log"
  "context"
)

type Client struct {
  postgresql.Client
}

func New(dsn string) Client{
  c, err := postgresql.NewClient(context.Context, dsn, 3)
if err!=nil{
  log.Fatal(err)
}
  return &Client{
    Client: c
  }
}