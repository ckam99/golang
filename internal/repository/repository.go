package repository

import (
	"github.com/ckam225/golang/fiber/internal/database"
	"gorm.io/gorm"
)

type Repository struct {
	DB   *gorm.DB
	Auth authRepository
	User userRepository
}

func NewRepositoy(cfg database.Config) *Repository {
	db := database.Init(&cfg, true)
	return &Repository{
		DB:   db,
		Auth: *NewAuthRepository(db),
		User: *NewUserRepository(db),
	}
}
