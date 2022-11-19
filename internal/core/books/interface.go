package books

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]Book, error)
	Find(ctx context.Context, id int) (Book, error)
	Store(ctx context.Context, book *Book) error
}

type UseCase interface {
	GetAll(ctx context.Context) ([]Book, error)
	Store(ctx context.Context, dto CreateDTO) (Book, error)
}
