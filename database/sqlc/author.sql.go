// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: author.sql

package database

import (
	"context"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (
    fullname, bio
) VALUES(
    $1,$2
) RETURNING id, fullname, bio, created_at, updated_at
`

type CreateAuthorParams struct {
	Fullname string `json:"fullname"`
	Bio      string `json:"bio"`
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, createAuthor, arg.Fullname, arg.Bio)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Fullname,
		&i.Bio,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE FROM authors WHERE id=$1
`

func (q *Queries) DeleteAuthor(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteAuthor, id)
	return err
}

const getAllAuthors = `-- name: GetAllAuthors :many
SELECT id, fullname, bio, created_at, updated_at FROM authors ORDER BY id LIMIT $1 OFFSET $2
`

type GetAllAuthorsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetAllAuthors(ctx context.Context, arg GetAllAuthorsParams) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, getAllAuthors, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Fullname,
			&i.Bio,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAuthor = `-- name: GetAuthor :one
SELECT id, fullname, bio, created_at, updated_at FROM authors WHERE id=$1 LIMIT 1
`

func (q *Queries) GetAuthor(ctx context.Context, id int32) (Author, error) {
	row := q.db.QueryRowContext(ctx, getAuthor, id)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Fullname,
		&i.Bio,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateAuthor = `-- name: UpdateAuthor :one
UPDATE authors SET fullname = $2, bio = $3, updated_at = now() WHERE id = $1 RETURNING id, fullname, bio, created_at, updated_at
`

type UpdateAuthorParams struct {
	ID       int32  `json:"id"`
	Fullname string `json:"fullname"`
	Bio      string `json:"bio"`
}

func (q *Queries) UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, updateAuthor, arg.ID, arg.Fullname, arg.Bio)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Fullname,
		&i.Bio,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
