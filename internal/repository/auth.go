package repository

import (
	"errors"

	"github.com/ckam225/golang/fiber/internal/entity"
	"github.com/ckam225/golang/fiber/internal/http/request"
	"github.com/ckam225/golang/fiber/internal/security"
	"github.com/ckam225/golang/fiber/internal/service"

	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) SignIn(rq *request.LoginRequest) (*entity.User, error) {
	user, err := service.GetUserByEmail(r.db, rq.Email)
	if err != nil {
		return nil, err
	}
	if !security.VerifyPassword(user.Password, rq.Password) {
		return nil, errors.New("bad credentials")
	}
	return user, nil
}

func (r *authRepository) SignUp(body request.RegisterRequest) (*entity.User, error) {
	user := entity.User{
		Name:  body.Name,
		Email: body.Email,
		Phone: body.Phone,
	}
	_, err := service.CreateUser(r.db, &user)
	return &user, err
}

func (r *authRepository) ChangeUserPassword(user *entity.User, newPassword string) error {
	var err error
	user.Password, err = security.HashPassword(newPassword)
	if err != nil {
		return err
	}
	err = r.db.Save(&user).Error
	return err
}
