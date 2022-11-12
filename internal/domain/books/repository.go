package books

import (
	"context"
	"database/sql"
)

type repo struct {
	*sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repo{
		DB: db,
	}
}

func (r *repo) GetAll(ctx context.Context, param *QueryFilterDTO) ([]Book, error) {
	q, args := getFilterQuery("select * from books", param)

	rows, err := r.QueryContext(ctx, q, args...)
	if err != nil {
		return []Book{}, err
	}

	books := []Book{}
	for rows.Next() {
		var book Book
		if err = rows.Scan(
			&book.ID,
			&book.Title,
			&book.Esbn,
			&book.Description,
			&book.AuthorID,
			&book.CreatedAt,
			&book.UpdatedAt,
		); err != nil {
			return []Book{}, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (r *repo) Find(ctx context.Context, id int64) (Book, error) {
	q := "select * from books where id=$1 limit 1"
	var book Book
	if err := r.QueryRowContext(ctx, q, id).
		Scan(
			&book.ID,
			&book.Title,
			&book.Esbn,
			&book.Description,
			&book.AuthorID,
			&book.CreatedAt,
			&book.UpdatedAt,
		); err != nil {
		return Book{}, err
	}
	return Book{}, nil
}

func (r *repo) Create(ctx context.Context, b *Book) error {
	q := `insert into books(
   title,esbn,description,author_id,updated_at
  ) values($1,$2,$3,$4,datetime('now') returning id,created_at,updated_at`
	if err := r.QueryRowContext(ctx, q, b.Title, b.Esbn, b.Description, b.AuthorID).
		Scan(&b.ID, &b.CreatedAt, &b.UpdatedAt); err != nil {
		return err
	}
	return nil
}

func (r *repo) Update(ctx context.Context, b *Book) error {
	q := `update books set
    title=coalesce($1,title),
    esbn=coalesce($2,esbn),
description=coalesce($3,description),
updated_at=datetime('now')
  returning id,updated_at`
	if err := r.QueryRowContext(ctx, q, b.Title, b.Esbn, b.Description).
		Scan(&b.ID, &b.UpdatedAt); err != nil {
		return err
	}
	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	panic("not implemented")
}

func (r *repo) Count(ctx context.Context, param *QueryFilterDTO) (int64, error) {
	q, args := getFilterQuery("select count(*) from books", param)
	var count int64
	if err := r.QueryRowContext(ctx, q, args...).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}
