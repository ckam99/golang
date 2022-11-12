package author

import (
	"app/pkg/clients/postgresql"
	"context"
	"fmt"
	"log"
	"strings"
)

type Repository interface {
	Create(ctx context.Context, author *Author) error
	GetAll(ctx context.Context, params FilterParamsDTO) (authors []Author, err error)
	Get(ctx context.Context, id int) (Author, error)
	Update(ctx context.Context, author *Author) error
	Delete(ctx context.Context, id int) error
	SoftDelete(ctx context.Context, author *Author) error
}

func NewRepository(client postgresql.Client) Repository {
	return &repository{
		client: client,
	}
}

type repository struct {
	client postgresql.Client
}

// Create implements Repository
func (r *repository) Create(ctx context.Context, author *Author) error {
	query := `
		INSERT INTO authors(name, bio) VALUES($1, $2) RETURNING id;
	   `
	if err := r.client.QueryRow(ctx, query, author.Name, author.Biography).Scan(&author.ID); err != nil {
		return postgresql.Error(err)
	}
	return nil
}

// Delete implements Repository
func (r *repository) Delete(ctx context.Context, id int) error {
	q := `DELETE FROM authors WHERE id=$1;`
	cmd, err := r.client.Exec(ctx, q, id)
	if err != nil {
		return postgresql.Error(err)
	}
	if !cmd.Delete() {
		return postgresql.Error(fmt.Errorf("can not delete author with id: %d", id))
	}
	return nil
}

// Get implements Repository
func (r *repository) Get(ctx context.Context, id int) (Author, error) {
	var author Author
	query := `
		SELECT * FROM authors WHERE id=$1 LIMIT 1;
	   `
	if err := r.client.QueryRow(ctx, query, id).Scan(
		&author.ID,
		&author.Name,
		&author.Biography,
		&author.CreatedAt,
		&author.UpdatedAt,
	); err != nil {
		return Author{}, postgresql.Error(err)
	}
	return author, nil
}

// GetAll implements Repository
func (r *repository) GetAll(ctx context.Context, params FilterParamsDTO) (authors []Author, err error) {
	q := `SELECT * FROM authors`
	args := []interface{}{}
	if len(params.OrderBy) > 0 {
		q += fmt.Sprintf(" ORDER BY %s", strings.Join(params.OrderBy, ","))
		if params.Ascending != "" {
			q += " " + params.Ascending
		}
	}
	if params.Limit > 0 {
		q += " LIMIT $1"
		args = append(args, params.Limit)
	}
	if params.Offset > 0 {
		q += fmt.Sprintf(" OFFSET $%d", len(args)+1)
		args = append(args, params.Limit)
	}
	rows, err := r.client.Query(ctx, q, args...)
	if err != nil {
		return []Author{}, postgresql.Error(err)
	}

	log.Println(q)
	log.Println(postgresql.DebugQuery(q, args))

	for rows.Next() {
		var author Author
		if err = rows.Scan(
			&author.ID,
			&author.Name,
			&author.Biography,
			&author.CreatedAt,
			&author.UpdatedAt,
		); err != nil {
			return []Author{}, postgresql.Error(err)
		}
		authors = append(authors, author)
	}
	return authors, nil
}

// Update implements Repository
func (r *repository) Update(ctx context.Context, author *Author) error {
	q := `
	UPDATE authors SET 
    name=COALESCE($1, name),
    bio=COALESCE($2,bio),
    updated_at=COALESCE($3, NOW())
	WHERE id=$4 RETURNING updated_at;
	`
	row := r.client.QueryRow(ctx, q, author.Name, author.Biography, author.UpdatedAt, author.ID)
	if err := row.Scan(&author.UpdatedAt); err != nil {
		return postgresql.Error(err)
	}
	return nil
}

// SoftDelete implements Repository
func (r *repository) SoftDelete(ctx context.Context, author *Author) error {
	q := `
	UPDATE authors SET
    deleted_at=COALESCE($1, NOW())
	WHERE id=$2 RETURNING deleted_at;
	`
	row := r.client.QueryRow(ctx, q, author.DeletedAt, author.ID)
	if err := row.Scan(&author.DeletedAt); err != nil {
		return postgresql.Error(err)
	}
	return nil
}
