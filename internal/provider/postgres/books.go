package postgres

import (
	"context"
	"main/internal/core/books"
)

type bookRepo struct {
	Client
}

func NewBookRepository(c Client) books.Repository {
	return &bookRepo{
		Client: c,
	}
}

func (r *bookRepo) GetAll(ctx context.Context) ([]books.Book, error) {
	panic("not implemented")
}

func (r *bookRepo) Find(ctx context.Context, id int) (books.Book, error) {
	panic("not implemented")
}

func (r *bookRepo) Store(ctx context.Context, book *books.Book) error {
	panic("not implemented")
}
