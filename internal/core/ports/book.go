package ports

import (
	"context"
	"example/grpc/internal/core/entity"
)

type BookRepository interface {
	GetAll(ctx context.Context, limit, offset int64) ([]entity.Book, error)
	// Count counts authors
	Count(ctx context.Context, limit, offset int64) (int64, error)
	Create(ctx context.Context, book *entity.Book) error
	Update(ctx context.Context, book *entity.Book) error
	Delete(ctx context.Context, bookID int64) error
	// GetByID get book by ID
	GetByID(ctx context.Context, bookID int64) (entity.Book, error)
}

type BookService interface {
	GetAll(ctx context.Context, limit, offset int64) ([]entity.Book, error)
	// Count counts authors
	Count(ctx context.Context, limit, offset int64) (int64, error)
	Create(ctx context.Context, book *entity.Book) error
	Update(ctx context.Context, book *entity.Book) error
	Delete(ctx context.Context, bookID int64) error
	// GetByID get book by ID
	GetByID(ctx context.Context, bookID int64) (entity.Book, error)
}
