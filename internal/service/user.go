package service

import (
	"github.com/ckam225/golang/echo/internal/database/postgres/storage"
	"github.com/ckam225/golang/echo/internal/entity"
	"github.com/google/uuid"
)

type IUserService interface {
	GetUsers(limit, offset int) ([]entity.User, error)
	GetUser(id uuid.UUID) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	CreateUser(t *entity.User) error
	UpdateUser(t *entity.User) error
	DeleteUser(id uuid.UUID) error
}

type userService struct {
	storage *storage.Store
}

func NewUserService(store *storage.Store) IUserService {
	u := userService{
		storage: store,
	}
	return &u
}

func (u *userService) GetUsers(limit int, offset int) ([]entity.User, error) {
	return u.storage.GetUsers(limit, offset)
}

func (u *userService) GetUser(id uuid.UUID) (entity.User, error) {
	return u.storage.GetUser(id)
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
	panic("not implemented") // TODO: Implement
}
