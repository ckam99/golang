package auth

import (
  "context"
)

type Repository interface {
   Find(ctx context.Context,user *User) error
   Create(ctx context.Context, user *User) error
   Update(ctx context.Context, user *User) error
   Delete(ctx context.Context, id int64, soft bool) error 
}

type Service interface {
   VerifyToken(token string) bool
   CreateToken(user User) string
   FindByID(ctx context.Context, id int64) (User, error)
   FindByEmail(ctx context.Context, email string) (User, error)
   Register(ctx context.Context, dto RegisterDTO)(User, error)
  Login(ctx context.Context, dto LoginDTO)(User, error)
}
