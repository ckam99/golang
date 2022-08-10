package entity

import "database/sql"

type Person struct {
	ID        int    `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

type Place struct {
	ID      int `db:"id"`
	Country string
	City    sql.NullString
	TelCode int
}
