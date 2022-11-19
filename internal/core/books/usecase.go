package books

import "context"

type useCase struct {
	repo Repository
}

func NewUseCase(r Repository) UseCase {
	return &useCase{
		repo: r,
	}
}

func (s *useCase) GetAll(ctx context.Context) ([]Book, error) {
	return s.repo.GetAll(ctx)
}

func (s *useCase) Store(ctx context.Context, dto CreateDTO) (Book, error) {
	b := Book{
		Title:       dto.Title,
		Description: dto.Description,
	}
	if err := s.repo.Store(ctx, &b); err != nil {
		return Book{}, err
	}
	return b, nil
}
