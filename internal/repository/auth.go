package repository

import (
	"errors"

	"github.com/ckam225/golang/fiber/internal/entity"
	"github.com/ckam225/golang/fiber/internal/http/request"
	"github.com/ckam225/golang/fiber/internal/security"
	"github.com/ckam225/golang/fiber/internal/service"

	"gorm.io/gorm"
)

type AuthRepository struct {
	Query *gorm.DB
}

func (r *AuthRepository) SignIn(rq *request.LoginRequest) (*entity.User, error) {
	user, err := service.GetUserByEmail(r.Query, rq.Email)
	if err != nil {
		return nil, err
	}
	if !security.VerifyPassword(user.Password, rq.Password) {
		return nil, errors.New("bad credentials")
	}
	return user, nil
}

func (r *AuthRepository) SignUp(body request.RegisterRequest) (*entity.User, error) {
	user := entity.User{
		Name:  body.Name,
		Email: body.Email,
		Phone: body.Phone,
	}
	_, err := service.CreateUser(r.Query, &user)
	return &user, err
}

func (r *AuthRepository) ChangeUserPassword(user *entity.User, newPassword string) error {
	var err error
	user.Password, err = security.HashPassword(newPassword)
	if err != nil {
		return err
	}
	err = r.Query.Save(&user).Error
	return err
}
