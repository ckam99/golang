package postgres

import (
	"context"
	"main/internal/domain"
)

type bookRepo struct {
	Client
}

func NewBookRepository(c Client) domain.BookRepository {
	return &bookRepo{
		Client: c,
	}
}

func (r *bookRepo) GetAll(ctx context.Context) ([]domain.Book, error) {
	panic("not implemented")
}

func (r *bookRepo) Find(ctx context.Context, id int) (domain.Book, error) {
	panic("not implemented")
}

func (r *bookRepo) Store(ctx context.Context, book *domain.Book) error {
	panic("not implemented")
}
