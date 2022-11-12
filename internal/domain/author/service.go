package author

import (
	"app/pkg/clients/postgresql"
	"context"
)

func NewService(client postgresql.Client) Service {
	return &service{
		repo: NewRepository(client),
	}
}

type service struct {
	repo Repository
}

// CreateAuthor implements Service
func (s *service) CreateAuthor(ctx context.Context, author *Author) error {
	return s.repo.Create(ctx, author)
}

// DeleteAuthor implements Service
func (s *service) DeleteAuthor(ctx context.Context, id int, isSoftDelete bool) error {
	if isSoftDelete {
		var auth Author
		return s.repo.SoftDelete(ctx, &auth)
	}
	return s.repo.Delete(ctx, id)
}

// SoftDeleteAuthor implements Service
func (s *service) SoftDeleteAuthor(ctx context.Context, author *Author) error {
	return s.repo.SoftDelete(ctx, author)
}

// UpdateAuthor implements Service
func (s *service) UpdateAuthor(ctx context.Context, author *Author) error {
	return s.repo.Update(ctx, author)
}

// FindAuthor implements Service
func (s *service) FindAuthor(ctx context.Context, id int) (Author, error) {
	return s.repo.Get(ctx, id)
}

// GetAuthors implements Service
func (s *service) GetAuthors(ctx context.Context, params FilterParamsDTO) ([]Author, error) {
	return s.repo.GetAll(ctx, params)
}
