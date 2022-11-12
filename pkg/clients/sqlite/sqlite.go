package sqlite

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)

func Connect(dsn string) *sql.DB{
  db, err := sql.Open("sqlite3",dsn)
  if err != nil{
    panic(err)
  }

  defer db.Close()

 // db.SetConnMaxLifetime(0)
//	db.SetMaxIdleConns(3)
//	db.SetMaxOpenConns(3)

  return db
}