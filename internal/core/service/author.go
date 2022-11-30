package service

import (
	"context"
	"example/grpc/internal/core/entity"
	"example/grpc/internal/core/ports"
)

type authorService struct {
	repo ports.AuthorRepository
}

// GetByID implements ports.AuthorService
func (s *authorService) GetByID(ctx context.Context, authorID int64) (entity.Author, error) {
	return s.repo.GetByID(ctx, authorID)
}

// Create implements ports.AuthorService
func (s *authorService) Create(ctx context.Context, author *entity.Author) error {
	return s.repo.Create(ctx, author)
}

// Delete implements ports.AuthorService
func (s *authorService) Delete(ctx context.Context, authorID int64) error {
	return s.repo.Delete(ctx, authorID)
}

// GetAll implements ports.AuthorService
func (s *authorService) GetAll(ctx context.Context) ([]entity.Author, error) {
	return s.repo.GetAll(ctx)
}

// Update implements ports.AuthorService
func (s *authorService) Update(ctx context.Context, author *entity.Author) error {
	return s.repo.Update(ctx, author)
}

func NewAuthorService(r ports.AuthorRepository) ports.AuthorService {
	return &authorService{
		repo: r,
	}
}
