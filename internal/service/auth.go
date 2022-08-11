package service

import (
	"errors"
	"time"

	"github.com/ckam225/golang/fiber/internal/entity"
	"github.com/ckam225/golang/fiber/internal/http/request"
	"github.com/ckam225/golang/fiber/internal/repository"
	"github.com/ckam225/golang/fiber/internal/security"
	"github.com/ckam225/golang/fiber/internal/utils"
	"github.com/ckam225/golang/fiber/pkg/mailer"
)

type IAuthService interface {
	SignIn(rq *request.LoginRequest) (*entity.User, error)
	SignUp(body request.RegisterRequest) (*entity.User, error)
	ChangeUserPassword(user *entity.User, newPassword string) error
	SendConfirmationEmail(user *entity.User) error
	VerifyConfirmationCode(req *request.EmailConfirmRequest) (*entity.Verifycode, error)
}

type authService struct {
	repo repository.Repository
}

func AuthService(repo repository.Repository) IAuthService {
	return &authService{
		repo: repo,
	}
}

func (s *authService) SignIn(rq *request.LoginRequest) (*entity.User, error) {
	user, err := s.repo.User.GetUserByEmail(rq.Email)
	if err != nil {
		return nil, err
	}
	if !security.VerifyPassword(user.Password, rq.Password) {
		return nil, errors.New("bad credentials")
	}
	return user, nil
}

func (s *authService) SignUp(body request.RegisterRequest) (*entity.User, error) {
	user := entity.User{
		Name:  body.Name,
		Email: body.Email,
		Phone: body.Phone,
	}
	_, err := s.repo.User.CreateUser(&user)
	return &user, err
}

func (s *authService) ChangeUserPassword(user *entity.User, newPassword string) error {
	var err error
	user.Password, err = security.HashPassword(newPassword)
	if err != nil {
		return err
	}
	err = s.repo.DB.Save(&user).Error
	return err
}

func (s *authService) SendConfirmationEmail(user *entity.User) error {
	vcode := entity.Verifycode{
		Email:  user.Email,
		Code:   utils.GenerateHashCode(),
		Target: "email",
	}

	if err := s.repo.DB.Create(&vcode).Error; err != nil {
		return err
	}
	mailer.NotificationChannel <- &mailer.Notification{
		To: []string{user.Email},
		Data: map[string]string{
			"email":  user.Email,
			"name":   user.Name,
			"code":   vcode.Code,
			"target": vcode.Target,
		},
		Subject:  "Registration",
		Template: "mail/confirm_action.tmpl",
	}
	return nil
}

func (s *authService) VerifyConfirmationCode(req *request.EmailConfirmRequest) (*entity.Verifycode, error) {
	var verifyCode entity.Verifycode
	if err := s.repo.DB.Where("email = ?", req.Email).Where("code = ?", req.Code).First(&verifyCode).Error; err != nil {
		return nil, err
	}
	duration := time.Since(verifyCode.CreatedAt)
	isExpired := false
	switch req.Target {
	case "email":
		isExpired = duration.Hours() > 24.0
	case "phone":
		isExpired = duration.Seconds() > 30.0
	}
	if isExpired {
		return nil, errors.New("token is expired")
	}
	return &verifyCode, nil
}
