package auth

import (
	"context"
)

type Repository interface {
	Find(ctx context.Context, user *User) error
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int64, soft bool) error
}

type Service interface {
	// Register: create new user
	Register(ctx context.Context, dto RegisterDTO) (User, error)
	// Login: authenticate user and generate token
	Login(ctx context.Context, dto LoginDTO) (Token, error)
	// Authenticate: verify user identity
	Authenticate(ctx context.Context, user *User, password string) error
	// RefreshAccessToken: create for user new  access token
	RefreshAccessToken(user *User, bearer string) (Token, error)
	// Get current authenticated user decrypting http authorization header
	GetCurrentUser(ctx context.Context, bearerToken string) (User, error)
	// FindByID: find user by id
	FindByID(ctx context.Context, id int64) (User, error)
	// FindByEmail: find user by email address
	FindByEmail(ctx context.Context, email string) (User, error)
	// FindByPhone: find user by phone number
	FindByPhone(ctx context.Context, phone string) (User, error)
}
