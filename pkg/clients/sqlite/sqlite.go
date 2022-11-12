package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func Connect(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	// defer db.Close()
	// db.SetConnMaxLifetime(0)
	//	db.SetMaxIdleConns(3)
	//	db.SetMaxOpenConns(3)
	return db, nil
}
