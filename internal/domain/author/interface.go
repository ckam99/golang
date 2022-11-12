package author

import "context"

type Repository interface {
	Create(ctx context.Context, author *Author) error
	GetAll(ctx context.Context, params FilterParamsDTO) (authors []Author, err error)
	Get(ctx context.Context, id int) (Author, error)
	Update(ctx context.Context, author *Author) error
	Delete(ctx context.Context, id int) error
	SoftDelete(ctx context.Context, author *Author) error
}

type Service interface {
	GetAuthors(ctx context.Context, params FilterParamsDTO) ([]Author, error)
	CreateAuthor(ctx context.Context, author *Author) error
	FindAuthor(ctx context.Context, id int) (Author, error)
	UpdateAuthor(ctx context.Context, author *Author) error
	DeleteAuthor(ctx context.Context, id int, isSoftDelete bool) error
	SoftDeleteAuthor(ctx context.Context, author *Author) error
}
