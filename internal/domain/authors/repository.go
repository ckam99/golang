package authors

import (
	"context"
	"database/sql"	
)

type repository struct {
	*sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		DB: db,
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
