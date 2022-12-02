package postgres

import (
	"context"
	"example/grpc/internal/core/entity"
	"example/grpc/internal/core/ports"
	"example/grpc/pkg/postgresql"
	"example/grpc/pkg/utils"
	"fmt"
)

type bookRepo struct {
	postgresql.Client
}

// Count implements ports.BookRepository
func (r *bookRepo) Count(ctx context.Context, limit, offset int64) (int64, error) {
	q := `select count(*) from authors`
	if limit > 0 {
		q += fmt.Sprint(` limit `, limit)
	}
	if offset >= 0 {
		q += fmt.Sprint(` offset `, offset)
	}
	var c int64
	if err := r.QueryRow(ctx, q).Scan(&c); err != nil {
		return 0, postgresql.Error(err)
	}
	return c, nil
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
	q := `update books set title=coalesce(nullif($2, ''),title), 
	description=coalesce(nullif($3, ''),description),
	published_at=coalesce($4,published_at),
	author_id=coalesce($5,author_id),
	updated_at=now() 
	where id=$1 returning updated_at`
	if err := r.QueryRow(ctx, q, book.ID, book.Title, book.Description, book.PublishedAt, book.AuthorID).
		Scan(&book.UpdatedAt); err != nil {
		return postgresql.Error(err)
	}
	return nil
}

// Create implements ports.BookRepository
func (r *bookRepo) Create(ctx context.Context, book *entity.Book) error {
	q := `insert into books(title, description, published_at, author_id) values($1,$2,$3, nullif($4, 0)) 
	returning id, created_at`
	if err := r.QueryRow(ctx, q, book.Title, book.Description, book.PublishedAt, book.AuthorID).
		Scan(&book.ID, &book.CreatedAt); err != nil {
		return postgresql.Error(err)
	}
	return nil
}

// GetAll implements ports.BookRepository
func (r *bookRepo) GetAll(ctx context.Context, limit, offset int64) ([]entity.Book, error) {
	q := `select * from books`
	if limit > 0 {
		q += fmt.Sprint(` limit `, limit)
	}
	if offset >= 0 {
		q += fmt.Sprint(` offset `, offset)
	}
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
