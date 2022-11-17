package books

import (
	"context"
	"main/pkg/clients/postgresql"
)

type repo struct {
	postgresql.Client
}

func NewRepository(db postgresql.Client) Repository {
	return &repo{
		Client: db,
	}
}

func (r *repo) GetAll(ctx context.Context, param *QueryFilterDTO) ([]Book, error) {
	q, args := getFilterQuery("select * from books", param)

	rows, err := r.Query(ctx, q, args...)
	if err != nil {
		return []Book{}, postgresql.Error(err)
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
			return []Book{}, postgresql.Error(err)
		}
		books = append(books, book)
	}
	return books, nil
}

func (r *repo) Find(ctx context.Context, id int64) (Book, error) {
	q := `select 
	books.id, books.title, books.esbn,books.description,books.created_at,books.updated_at, 
	row_to_json(authors.*) author
	from books left join authors on authors.id = books.author_id
	where books.id=$1 limit 1`
	var book Book
	if err := r.QueryRow(ctx, q, id).
		Scan(
			&book.ID,
			&book.Title,
			&book.Esbn,
			&book.Description,
			// &book.AuthorID,
			&book.CreatedAt,
			&book.UpdatedAt,
			&book.Author,
		); err != nil {

		return Book{}, postgresql.Error(err)

	}
	return book, nil
}

func (r *repo) Create(ctx context.Context, b *Book) error {
	q := `insert into books(
   title,esbn,description,author_id,updated_at
  ) values($1,$2,$3,$4,now()) returning id,created_at,updated_at`
	if err := r.QueryRow(ctx, q, b.Title, b.Esbn, b.Description, b.AuthorID).
		Scan(&b.ID, &b.CreatedAt, &b.UpdatedAt); err != nil {
		return postgresql.Error(err)
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
	if err := r.QueryRow(ctx, q, b.Title, b.Esbn, b.Description).
		Scan(&b.ID, &b.UpdatedAt); err != nil {
		return postgresql.Error(err)
	}
	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	panic("not implemented")
}

func (r *repo) Count(ctx context.Context, param *QueryFilterDTO) (int64, error) {
	q, args := getFilterQuery("select count(*) from books", param)
	var count int64
	if err := r.QueryRow(ctx, q, args...).Scan(&count); err != nil {
		return 0, postgresql.Error(err)
	}
	return count, nil
}
