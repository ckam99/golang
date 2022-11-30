package ports

import (
	"context"
	"example/grpc/internal/core/entity"
)

type AuthorRepository interface {
	// GetAll get list authors
	GetAll(ctx context.Context) ([]entity.Author, error)
	// Create create new autthor
	Create(ctx context.Context, author *entity.Author) error
	// Update update author by id
	Update(ctx context.Context, author *entity.Author) error
	// Delete delete author from the database
	Delete(ctx context.Context, authorID int64) error
	// GetByID get author by ID
	GetByID(ctx context.Context, authorID int64) (entity.Author, error)
}

type AuthorService interface {
	// GetAll get list authors
	GetAll(ctx context.Context) ([]entity.Author, error)
	// Create create new autthor
	Create(ctx context.Context, author *entity.Author) error
	// Update update author by id
	Update(ctx context.Context, author *entity.Author) error
	// Delete delete author from the database
	Delete(ctx context.Context, authorID int64) error
	// GetByID get author by ID
	GetByID(ctx context.Context, authorID int64) (entity.Author, error)
}
