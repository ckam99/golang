package entity

import (
	"database/sql"
	"time"
)

type Person struct {
	ID        int       `json:"id" db:"id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Place struct {
	ID      int `db:"id"`
	Country string
	City    sql.NullString
	TelCode int
}
