package postgres

import (
	"context"
	"main/internal/core/authors"
)

type authorRepo struct {
	Client
}

func NewAuthorRepository(c Client) authors.Repository {
	return &authorRepo{
		Client: c,
	}
}

func (r *authorRepo) GetAll(ctx context.Context) ([]authors.Author, error) {
	panic("not implemented")
}

func (r *authorRepo) Find(ctx context.Context, id int) (authors.Author, error) {
	panic("not implemented")
}

func (r *authorRepo) Store(ctx context.Context, author *authors.Author) error {
	panic("not implemented")
}
