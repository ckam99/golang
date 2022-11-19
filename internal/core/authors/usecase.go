package authors

import "context"

type useCase struct {
	repo Repository
}

func NewUseCase(r Repository) UseCase {
	return &useCase{
		repo: r,
	}
}

func (s *useCase) GetAll(ctx context.Context) ([]Author, error) {
	return s.repo.GetAll(ctx)
}

func (s *useCase) Store(ctx context.Context, dto CreateDTO) (Author, error) {
	b := Author{
		Name: dto.Name,
		Bio:  dto.Bio,
	}
	if err := s.repo.Store(ctx, &b); err != nil {
		return Author{}, err
	}
	return b, nil
}
