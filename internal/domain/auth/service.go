package auth

import (
	"context"
	"main/pkg/clients/postgresql"
	"main/pkg/utils"

	"github.com/golang-jwt/jwt/v4"
)

type service struct {
	repo Repository
}

func NewService(client postgresql.Client) Service {
	return &service{
		repo: NewRepository(client),
	}
}

func (s *service) FindByID(ctx context.Context, id int64) (User, error) {
	user := User{
		ID: id,
	}
	if err := s.repo.Find(ctx, &user); err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *service) FindByEmail(ctx context.Context, email string) (User, error) {
	user := User{
		Email: email,
	}
	if err := s.repo.Find(ctx, &user); err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *service) FindByPhone(ctx context.Context, phone string) (User, error) {
	user := User{
		Phone: &phone,
	}
	if err := s.repo.Find(ctx, &user); err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *service) Register(ctx context.Context, dto RegisterDTO) (User, error) {
	hash, err := HashPassword(dto.Password)
	if err != nil {
		return User{}, err
	}
	user := User{
		Email:    dto.Email,
		FullName: dto.FullName,
		Password: &hash,
		Role:     "user",
	}
	if dto.Phone != "" {
		user.Phone = &dto.Phone
	}
	if err := s.repo.Create(ctx, &user); err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *service) Authenticate(ctx context.Context, user *User, password string) error {
	if err := s.repo.Find(ctx, user); err != nil {
		return err
	}
	if !VerifyPassword(*user.Password, password) {
		return utils.ErrInvalidCredentials
	}
	return nil
}

func (s *service) Login(ctx context.Context, dto LoginDTO) (Token, error) {
	user := User{
		Email: dto.Email,
	}
	if dto.Phone != "" {
		user.Phone = &dto.Phone
	}
	if err := s.Authenticate(ctx, &user, dto.Password); err != nil {
		return Token{}, err
	}
	claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   TOKEN_EXPIRE_TIME,
	}
	token, err := GenerateToken(claims)
	if err != nil {
		return Token{}, err
	}
	return Token{
		ID:           user.ID,
		Email:        user.Email,
		AccessToken:  token,
		RefreshToken: token,
	}, nil
}

func (s *service) RefreshAccessToken(user *User, bearer string) (Token, error) {
	oldToken, err := DecodeToken(bearer)
	if err != nil {
		return Token{}, err
	}
	claims := jwt.MapClaims{
		"email": user.Email,
		"token": oldToken,
		"exp":   TOKEN_EXPIRE_TIME,
	}
	newToken, err := GenerateToken(claims)
	if err != nil {
		return Token{}, err
	}
	return Token{
		ID:           user.ID,
		Email:        user.Email,
		AccessToken:  oldToken,
		RefreshToken: newToken,
	}, nil
}

func (s *service) GetCurrentUser(ctx context.Context, bearerToken string) (User, error) {
	token, err := VerifyToken(bearerToken)
	if err != nil {
		return User{}, err
	}
	claims := token.Claims.(jwt.MapClaims)
	user := User{
		Email: claims["email"].(string),
	}
	if err = s.repo.Find(ctx, &user); err != nil {
		return User{}, err
	}
	return user, nil
}
