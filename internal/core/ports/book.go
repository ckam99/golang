package ports

import (
	"context"
	"example/grpc/internal/core/entity"
)

type BookRepository interface {
	GetAll(ctx context.Context) ([]entity.Book, error)
	Create(ctx context.Context, book *entity.Book) error
}

type BookService interface {
	GetAll(ctx context.Context) ([]entity.Book, error)
	Create(ctx context.Context, book *entity.Book) error
}
