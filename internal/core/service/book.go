package service

import (
	"context"
	"example/grpc/internal/core/entity"
	"example/grpc/internal/core/ports"
)

type bookService struct {
	repo ports.BookRepository
}

// Delete implements ports.BookService
func (*bookService) Delete(ctx context.Context, bookID int64) error {
	panic("unimplemented")
}

// Update implements ports.BookService
func (*bookService) Update(ctx context.Context, book *entity.Book) error {
	panic("unimplemented")
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
