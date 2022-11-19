package authors

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]Author, error)
	Find(ctx context.Context, id int) (Author, error)
	Store(ctx context.Context, book *Author) error
}

type UseCase interface {
	GetAll(ctx context.Context) ([]Author, error)
	Store(ctx context.Context, dto CreateDTO) (Author, error)
}
