package database

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	DBDRIVER = "postgres"
	DBURL    = "postgres://postgres:postgres@host.docker.internal:54323/demo?sslmode=disable"
)

var store *Store

func TestMain(m *testing.M) {
	cnx, err := sql.Open(DBDRIVER, DBURL)
	if err != nil {
		log.Fatal("connot connect to the database: ", err.Error())
	}
	store = NewStore(cnx)
	os.Exit(m.Run())
}
