package service

import (
	"github.com/ckam225/golang/fiber/internal/config"
	"github.com/ckam225/golang/fiber/internal/repository"
	"gorm.io/gorm"
)

var repo repository.Repository

type Service struct {
	IUserService
	IAuthService
}

func NewService(cfg config.Configuration) *Service {
	repo = *repository.NewRepositoy(*cfg.Database)
	return &Service{
		UserService(repo),
		AuthService(repo),
	}
}

func (s *Service) GetDB() *gorm.DB {
	return repo.DB
}
