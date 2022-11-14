package authors

import (
	"context"
	"main/pkg/clients/postgresql"
)

type service struct {
	repo Repository
}

func NewService(db postgresql.Client) Service {
	return &service{
		repo: NewRepository(db),
	}
}

func (s *service) GetAll(ctx context.Context, param *QueryFilterDTO) ([]Author, error) {
	panic("not implemented")
}

func (s *service) Find(ctx context.Context, id int64) (Author, error) {
	panic("not implemented")
}

func (s *service) Create(ctx context.Context, dto CreateDTO) error {
	panic("not implemented")
}

func (s *service) Update(ctx context.Context, dto UpdateDTO) error {
	panic("not implemented")
}

func (s *service) Delete(ctx context.Context, id int64) error {
	panic("not implemented")
}
