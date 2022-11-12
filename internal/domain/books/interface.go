package books

import (
	"context"
)

type Repository interface {
	GetAll(ctx context.Context, param *QueryFilterDTO) ([]Book, error)
	Create(ctx context.Context, book *Book) error
	Find(ctx context.Context, id int64) (Book, error)
	Update(ctx context.Context, book *Book) error
	Delete(ctx context.Context, id int64) error
	Count(ctx context.Context, param *QueryFilterDTO) (int64, error)
}

type Service interface {
	GetAll(ctx context.Context, param *QueryFilterDTO) ([]Book, error)
	Create(ctx context.Context, book CreateDTO) (Book, error)
	Find(ctx context.Context, id int64) (Book, error)
	Update(ctx context.Context, id int64, book UpdateDTO) (Book, error)
	Delete(ctx context.Context, id int64) error
}
