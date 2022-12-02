package service

import (
	"context"
	"example/grpc/internal/core/entity"
	"example/grpc/internal/core/ports"
)

type bookService struct {
	repo ports.BookRepository
}

// Count implements ports.BookService
func (s *bookService) Count(ctx context.Context, limit, offset int64) (int64, error) {
	return s.repo.Count(ctx, limit, offset)
}

// GetByID implements ports.BookService
func (s *bookService) GetByID(ctx context.Context, bookID int64) (entity.Book, error) {
	return s.repo.GetByID(ctx, bookID)
}

// Delete implements ports.BookService
func (s *bookService) Delete(ctx context.Context, bookID int64) error {
	return s.repo.Delete(ctx, bookID)
}

// Update implements ports.BookService
func (s *bookService) Update(ctx context.Context, book *entity.Book) error {
	return s.repo.Update(ctx, book)
}

// Create implements ports.BookService
func (s *bookService) Create(ctx context.Context, book *entity.Book) error {
	return s.repo.Create(ctx, book)
}

// GetAll implements ports.BookService
func (s *bookService) GetAll(ctx context.Context, limit, offset int64) ([]entity.Book, error) {
	return s.repo.GetAll(ctx, limit, offset)
}

func NewBookService(r ports.BookRepository) ports.BookService {
	return &bookService{
		repo: r,
	}
}
