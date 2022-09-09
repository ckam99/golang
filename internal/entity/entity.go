package entity

import "database/sql"

type Person struct {
	ID        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email"`
}

type Place struct {
	ID      int `db:"id"`
	Country string
	City    sql.NullString
	TelCode int
}
