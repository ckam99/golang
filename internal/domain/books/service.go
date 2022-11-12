package books

import (
	"context"
	"database/sql"
)

type service struct {
	repo Repository
}

func NewService(db *sql.DB) Service {
	return &service{
		repo: NewRepository(db),
	}
}

func (s *service) GetAll(ctx context.Context, param *QueryFilterDTO) ([]Book, error) {
	return s.repo.GetAll(ctx, param)
}

func (s *service) Create(ctx context.Context, dto CreateDTO) (Book, error) {
	book := Book{
		Title:       dto.Title,
		Esbn:        &dto.Esbn,
		Description: &dto.Description,
		AuthorID:    dto.AuthorID,
	}
	if err := s.repo.Create(ctx, &book); err != nil {
		return Book{}, err
	}
	return book, nil
}

func (s *service) Find(ctx context.Context, id int64) (Book, error) {
	return s.repo.Find(ctx, id)
}

func (s *service) Update(ctx context.Context, id int64, dto UpdateDTO) (Book, error) {
	book := Book{
		ID:          id,
		Title:       dto.Title,
		Esbn:        &dto.Esbn,
		Description: &dto.Description,
		AuthorID:    dto.AuthorID,
	}
	if err := s.repo.Update(ctx, &book); err != nil {
		return Book{}, err
	}
	return book, nil
}

func (s *service) Delete(ctx context.Context, id int64) error {
	panic("not implemented")
}
