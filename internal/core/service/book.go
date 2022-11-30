package service

import (
	"context"
	"example/grpc/internal/core/entity"
	"example/grpc/internal/core/ports"
)

type bookService struct {
	repo ports.BookRepository
}

// Create implements ports.BookService
func (s *bookService) Create(ctx context.Context, book *entity.Book) error {
	return s.repo.Create(ctx, book)
}

// GetAll implements ports.BookService
func (s *bookService) GetAll(ctx context.Context) ([]entity.Book, error) {
	return s.repo.GetAll(ctx)
}

func NewBookService(r ports.BookRepository) ports.BookService {
	return &bookService{
		repo: r,
	}
}
