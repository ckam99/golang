package main

import (
	"main/pkg/migrator"

	_ "github.com/lib/pq"
)

func main() {
	// postgres://postgres:postgres@host.docker.internal/demo?sslmode=disable
	migrator.CommandLine()
}
