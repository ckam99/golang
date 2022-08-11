package repository

import (
	"github.com/ckam225/golang/fiber/internal/database"
	"gorm.io/gorm"
)

type Repository struct {
	DB   *gorm.DB
	User IUserRepository
}

func NewRepositoy(cfg database.Config) *Repository {
	db := database.Init(&cfg, true)
	return &Repository{
		DB:   db,
		User: UserRepository(db),
	}
}
