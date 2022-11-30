package postgres

import (
	"context"
	"example/grpc/internal/core/entity"
	"example/grpc/internal/core/ports"
	"example/grpc/pkg/postgresql"
	"example/grpc/pkg/utils"
)

type authorRepo struct {
	postgresql.Client
}

// GetByID implements ports.AuthorRepository
func (r *authorRepo) GetByID(ctx context.Context, authorID int64) (entity.Author, error) {
	q := `select * from authors where id=$1 limit 1`
	var author entity.Author
	if err := r.QueryRow(ctx, q, authorID).Scan(
		&author.ID,
		&author.Name,
		&author.Biography,
		&author.CreatedAt,
		&author.UpdatedAt,
	); err != nil {
		return entity.Author{}, postgresql.Error(err)
	}
	return author, nil
}

// Create implements ports.AuthorRepository
func (r *authorRepo) Create(ctx context.Context, author *entity.Author) error {
	q := `insert into authors(name, biography) values($1,$2) 
	returning id, created_at`
	if err := r.QueryRow(ctx, q, author.Name, author.Biography).
		Scan(&author.ID, &author.CreatedAt); err != nil {
		return postgresql.Error(err)
	}
	return nil
}

// Delete implements ports.AuthorRepository
func (r *authorRepo) Delete(ctx context.Context, authorID int64) error {
	q := `delete from authors where id=$1;`
	cmd, err := r.Exec(ctx, q, authorID)
	if err != nil {
		return postgresql.Error(err)
	}
	if cmd.RowsAffected() == 0 {
		return utils.ErrNoEntity
	}
	return nil
}

// GetAll implements ports.AuthorRepository
func (r *authorRepo) GetAll(ctx context.Context) ([]entity.Author, error) {
	q := `select * from authors`
	rows, err := r.Query(ctx, q)
	if err != nil {
		return []entity.Author{}, postgresql.Error(err)
	}
	result := make([]entity.Author, 0)
	for rows.Next() {
		var author entity.Author
		if err := rows.Scan(
			&author.ID,
			&author.Name,
			&author.Biography,
			&author.CreatedAt,
			&author.UpdatedAt,
		); err != nil {
			return []entity.Author{}, postgresql.Error(err)
		}
		result = append(result, author)
	}
	return result, nil
}

// Update implements ports.AuthorRepository
func (r *authorRepo) Update(ctx context.Context, author *entity.Author) error {
	q := `update authors set name=coalesce(nullif($1, ''),name), 
	biography=coalesce(nullif($2, ''),biography), 
	updated_at=now() where id=$3 returning updated_at`
	if err := r.QueryRow(ctx, q, author.Name, author.Biography, author.ID).
		Scan(&author.UpdatedAt); err != nil {
		return postgresql.Error(err)
	}
	return nil
}

func NewAuthorRepository(c postgresql.Client) ports.AuthorRepository {
	return &authorRepo{
		Client: c,
	}
}
