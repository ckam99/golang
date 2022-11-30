package ports

import (
	"context"
	"example/grpc/internal/core/entity"
)

type BookRepository interface {
	GetAll(ctx context.Context) ([]entity.Book, error)
	Create(ctx context.Context, book *entity.Book) error
	Update(ctx context.Context, book *entity.Book) error
	Delete(ctx context.Context, bookID int64) error
}

type BookService interface {
	GetAll(ctx context.Context) ([]entity.Book, error)
	Create(ctx context.Context, book *entity.Book) error
	Update(ctx context.Context, book *entity.Book) error
	Delete(ctx context.Context, bookID int64) error
}
