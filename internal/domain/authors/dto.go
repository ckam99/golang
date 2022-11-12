package authors

type QueryFilterDTO struct {
	FullName string `json:"full_name" validate:"empty=true" query:"full_name"`
	Limit    *int64 `json:"limit" validate:"nil=true | gt=0"`
	Offset   *int64 `json:"offset" validate:"nil=true | gte=0"`
	OrderBy  string `json:"order_by" validate:"empty=true|one_of=id,full_name,created_at" query:"order_by"`
	Sort     string `json:"sort" validate:"empty=true | one_of=asc,desc"`
}

type CreateDTO struct {
	Fullname  string  `json:"full_name" validate:"empty=false & min=2"`
	Biography *string `json:"author_id"`
}

type UpdateDTO struct {
	Fullname  string  `json:"full_name" validate:"empty=false & min=2"`
	Biography *string `json:"author_id" validate:"nil=true | min=10"`
}
