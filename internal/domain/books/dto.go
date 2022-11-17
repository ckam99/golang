package books

type CreateDTO struct {
	Title       string `json:"title" validate:"empty=false & gte=2"`
	AuthorID    int64  `json:"author_id" validate:"gt=0"`
	Esbn        string `json:"esbn"`
	Description string `json:"description"`
}

type UpdateDTO struct {
	Title       string `json:"title" validate:"empty=false & gte=2"`
	AuthorID    int64  `json:"author_id" validate:"gt=0"`
	Esbn        string `json:"esbn"`
	Description string `json:"description"`
}

type QueryFilterDTO struct {
	Title    string `json:"title" validate:"empty=true"`
	AuthorID *int64 `json:"author_id" validate:"nil=true | gt=0" query:"author_id"`
	Limit    *int64 `json:"limit" validate:"nil=true | gt=0"`
	Offset   *int64 `json:"offset" validate:"nil=true | gte=0"`
	OrderBy  string `json:"order_by" validate:"empty=true|one_of=id,author_id,title" query:"order_by"`
	Sort     string `json:"sort" validate:"empty=true | one_of=asc,desc"`
}
