package author

import "database/sql"

type Author struct {
	ID        int           `json:"id" `
	Name      string        `json:"name" `
	Biography string        `json:"biography" `
	UpdatedAt *sql.NullTime `json:"updated_at,omitempty"`
	CreatedAt *sql.NullTime `json:"created_at,omitempty"`
	DeletedAt *sql.NullTime `json:"deleted_at,omitempty"`
}
