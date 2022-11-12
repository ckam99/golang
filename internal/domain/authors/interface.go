package authors

import "context"

type Repository interface {
	GetAll(ctx context.Context, param *QueryFilterDTO) ([]Author, error)
	Count(ctx context.Context, param *QueryFilterDTO) (int64, error)
	Find(ctx context.Context, id int64) (Author, error)
	Create(ctx context.Context, author *Author) error
	Update(ctx context.Context, author *Author) error
	Delete(ctx context.Context, id int64) error
}

type Service interface {
	GetAll(ctx context.Context, param *QueryFilterDTO) ([]Author, error)
	Find(ctx context.Context, id int64) (Author, error)
	Create(ctx context.Context, dto CreateDTO) error
	Update(ctx context.Context, dto UpdateDTO) error
	Delete(ctx context.Context, id int64) error
}
