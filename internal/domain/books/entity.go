package books

import "time"

type Book struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Esbn        *string    `json:"esbn"`
	AuthorID    *int64     `json:"author_id,omitempty"`
	Author      *Author    `json:"author,omitempty"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type Author struct {
	ID        int64   `json:"id"`
	FullName  string  `json:"full_name"`
	Biography *string `json:"biography"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}
