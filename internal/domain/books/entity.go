package books

type Book struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Esbn        *string `json:"esbn"`
	AuthorID    string  `json:"author_id,omitempty"`
	Author      *Author `json:"author"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

type Author struct {
	ID        int64   `json:"id"`
	FullName  string  `json:"full_name"`
	Biography *string `json:"author_id"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}
