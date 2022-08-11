package service

import (
	"github.com/ckam225/golang/fiber/internal/config"
	"github.com/ckam225/golang/fiber/internal/repository"
)

type Service struct {
	User IUserService
	Auth IAuthService
	Repo repository.Repository
}

func NewService(cfg config.Configuration) *Service {
	repo := repository.NewRepositoy(*cfg.Database)
	return &Service{
		User: UserService(*repo),
		Auth: AuthService(*repo),
		Repo: *repo,
	}
}
