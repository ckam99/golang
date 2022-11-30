package postgres

import (
	"context"
	"example/grpc/internal/core/entity"
	"example/grpc/internal/core/ports"
	"example/grpc/pkg/postgresql"
	"example/grpc/pkg/utils"
)

type bookRepo struct {
	postgresql.Client
}

// GetByID implements ports.AuthorRepository
func (r *bookRepo) GetByID(ctx context.Context, bookID int64) (entity.Book, error) {
	q := `select * from books where id=$1 limit 1`
	var book entity.Book
	if err := r.QueryRow(ctx, q, bookID).Scan(
		&book.ID,
		&book.Title,
		&book.Description,
		&book.PublishedAt,
		&book.CreatedAt,
		&book.UpdatedAt,
		&book.AuthorID,
	); err != nil {
		return entity.Book{}, postgresql.Error(err)
	}
	return book, nil
}

// Delete implements ports.BookRepository
func (r *bookRepo) Delete(ctx context.Context, bookID int64) error {
	q := `delete from books where id=$1;`
	cmd, err := r.Exec(ctx, q, bookID)
	if err != nil {
		return postgresql.Error(err)
	}
	if cmd.RowsAffected() == 0 {
		return utils.ErrNoEntity
	}
	return nil
}

// Update implements ports.BookRepository
func (r *bookRepo) Update(ctx context.Context, book *entity.Book) error {
	q := `update books set title=coalesce(nullif($1, ''),title), 
	description=coalesce(nullif($2, ''),description),
	published_at=coalesce($3,published_at),
	updated_at=now() 
	where id=$4 returning updated_at`
	if err := r.QueryRow(ctx, q, book.Title, book.Description, book.PublishedAt, book.ID).
		Scan(&book.UpdatedAt); err != nil {
		return postgresql.Error(err)
	}
	return nil
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
			&book.AuthorID,
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
