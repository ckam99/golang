# Golang migrate

```go
package main

import (
	"github.com/ckam225/go-migrate"
	"log"
	_ "github.com/lib/pq"
)

func main() {
	m, err := migrate.New("./migrations", "postgres",
		"postgres://postgres:postgres@localhost/demo?sslmode=disable",
		&migrate.Config{},
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Migrate(); err != nil {
		panic(err)
	}
}

```