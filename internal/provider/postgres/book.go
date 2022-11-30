package postgres

import (
	"context"
	"example/grpc/internal/core/entity"
	"example/grpc/internal/core/ports"
	"example/grpc/pkg/postgresql"
)

type bookRepo struct {
	postgresql.Client
}

// Create implements ports.BookRepository
func (r *bookRepo) Create(ctx context.Context, book *entity.Book) error {
	q := `insert into books(title, description, published_at) values($1,$2,$3) 
	returning id, created_at`
	if err := r.QueryRow(ctx, q, book.Title, book.Description, book.PublishedAt).
		Scan(&book.ID, &book.CreatedAt); err != nil {
		return postgresql.Error(err)
	}
	return nil
}

// GetAll implements ports.BookRepository
func (r *bookRepo) GetAll(ctx context.Context) ([]entity.Book, error) {
	q := `select * from books`
	rows, err := r.Query(ctx, q)
	if err != nil {
		return []entity.Book{}, postgresql.Error(err)
	}
	result := make([]entity.Book, 0)
	for rows.Next() {
		var book entity.Book
		if err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.PublishedAt,
			&book.CreatedAt,
			&book.UpdatedAt,
		); err != nil {
			return []entity.Book{}, postgresql.Error(err)
		}
		result = append(result, book)
	}
	return result, nil
}

func NewBookRepository(c postgresql.Client) ports.BookRepository {
	return &bookRepo{
		Client: c,
	}
}
