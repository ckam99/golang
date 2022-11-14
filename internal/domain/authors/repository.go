package authors

import (
	"context"
	"main/pkg/clients/postgresql"
)

type repository struct {
	postgresql.Client
}

func NewRepository(db postgresql.Client) Repository {
	return &repository{
		Client: db,
	}
}

func (r *repository) GetAll(ctx context.Context, param *QueryFilterDTO) ([]Author, error) {
	panic("not implemented")
}

func (r *repository) Count(ctx context.Context, param *QueryFilterDTO) (int64, error) {
	panic("not implemented")
}

func (r *repository) Find(ctx context.Context, id int64) (Author, error) {
	panic("not implemented")
}

func (r *repository) Create(ctx context.Context, author *Author) error {
	panic("not implemented")
}

func (r *repository) Update(ctx context.Context, author *Author) error {
	panic("not implemented")
}

func (r *repository) Delete(ctx context.Context, id int64) error {
	panic("not implemented")
}
