package service

import (
	"fmt"

	"github.com/ckam225/golang/fiber-sqlx/internal/database/postgres/storage"
	"github.com/ckam225/golang/fiber-sqlx/internal/entity"
	"github.com/google/uuid"
)

type IUserService interface {
	GetUsers(limit, offset int) ([]entity.User, error)
	FindUser(id uuid.UUID) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	CreateUser(t *entity.User) error
	UpdateUser(t *entity.User) error
	DeleteUser(id uuid.UUID) error
	IsEmailExist(email string) bool
}

type userService struct {
	storage *storage.Store
}

func UserService(store *storage.Store) IUserService {
	u := userService{
		storage: store,
	}
	return &u
}

func (u *userService) GetUsers(limit int, offset int) ([]entity.User, error) {
	return u.storage.GetUsers(limit, offset)
}

func (u *userService) FindUser(id uuid.UUID) (entity.User, error) {
	return u.storage.FindUser(id)
}

func (u *userService) CreateUser(t *entity.User) error {
	return u.storage.CreateUser(t)
}

func (u *userService) UpdateUser(t *entity.User) error {
	return u.storage.UpdateUser(t)
}

func (u *userService) DeleteUser(id uuid.UUID) error {
	return u.DeleteUser(id)
}

func (u *userService) GetUserByEmail(email string) (entity.User, error) {
	return u.storage.FindUserBy("email", email)
}

func (u *userService) IsEmailExist(email string) bool {
	count, err := u.storage.CountUserBy("email", email)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return count > 0
}
