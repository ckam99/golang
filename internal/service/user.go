package service

import (
	"github.com/ckam225/golang/fiber/internal/entity"
	"github.com/ckam225/golang/fiber/internal/http/request"
	"github.com/ckam225/golang/fiber/internal/repository"
)

type userService struct {
	repo repository.Repository
}

type IUserService interface {
	GetAllUsers(p request.UserFilterParam) (*[]entity.User, error)
	GetUser(user *entity.User) (*entity.User, error)
	GetUserByID(id int) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	CreateUser(obj *request.CreateUser) (*entity.User, error)
	UpdateUser(userId int, payload *request.UpdateUser) (*entity.User, error)
	DeleteUser(userId uint, isSoftDelete bool) error
}

func UserService(repo repository.Repository) IUserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetAllUsers(p request.UserFilterParam) (*[]entity.User, error) {
	users, err := s.repo.User.GetAllUsers(p.Limit, p.Skip)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) CreateUser(obj *request.CreateUser) (*entity.User, error) {
	user := entity.User{
		Name:  obj.Name,
		Email: obj.Email,
		Phone: obj.Phone,
	}
	_, err := s.repo.User.CreateUser(&user)
	return &user, err
}

func (s *userService) GetUser(user *entity.User) (*entity.User, error) {
	return s.repo.User.GetUser(user)
}

func (s *userService) GetUserByID(id int) (*entity.User, error) {
	return s.repo.User.GetUserByID(id)
}

func (s *userService) GetUserByEmail(email string) (*entity.User, error) {
	return s.repo.User.GetUserByEmail(email)
}

func (s *userService) UpdateUser(userId int, payload *request.UpdateUser) (*entity.User, error) {
	return s.repo.User.UpdateUser(&entity.User{
		ID:    uint(userId),
		Name:  payload.Name,
		Phone: payload.Phone,
	})
}

func (s *userService) DeleteUser(userId uint, isSoftDelete bool) error {
	return s.repo.User.DeleteUser(userId, isSoftDelete)
}
