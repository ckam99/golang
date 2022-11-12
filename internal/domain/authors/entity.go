package authors

type Author struct {
	ID        string  `json:"id"`
	Fullname  string  `json:"full_name"`
	Biography *string `json:"author_id"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}
