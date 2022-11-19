package domain

import "context"

type service struct {
	repo BookRepository
}

func NewBookService(r BookRepository) BookService {
	return &service{
		repo: r,
	}
}

func (s *service) GetAll(ctx context.Context) ([]Book, error) {
	return s.repo.GetAll(ctx)
}

func (s *service) Store(ctx context.Context, dto CreateBookDTO) (Book, error) {
	b := Book{
		Title:       dto.Title,
		Description: dto.Description,
	}
	if err := s.repo.Store(ctx, &b); err != nil {
		return Book{}, err
	}
	return b, nil
}
