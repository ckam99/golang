package auth

import (
  "time"
  
)
type service struct {
  repo Repository
}

func NewService(client postgresql.Client) Service{
  return &service{
    repo: NewRepository(client)
  }
}

func(s *service) RefreshAccessToken(user *User, bearer string) (token, error) {
  oldToken, err := DecodeToken(bearer)
	if err != nil {
		return Token{}, err
  }
	claims := jwt.MapClaims{
		"email": user.Email,
		"token": oldToken,
		"exp":   time.Now().Add(time.Minute * 2).Unix(),
	}
	newToken, err:= GenerateToken(claims)
  if err != nil {
		return Token{}, err
  }

return Token{
		ID:           user.ID,
		Email:        user.Email,
		AccessToken:  currentToken,
		RefreshToken: newToken,
}, nil
}

func(s *service) GetAuthUser(bearerToken string)(User,error){
 token, err := VerifyToken(beareToken)
	if err != nil {
		return User{}, err
	}
	claims := token.Claims.(jwt.MapClaims)
  user := entity.User{
		Email: claims["email"].(string),
  }
	if err = s.repo.Find(ctx,&user); err != nil {
		return User{}, err
  }
  return user, nil
}

func(s *service) FindByID(ctx context.Context, id int64) (User, error){
 panic("not implemented")
}

func(s *service) FindByEmail(ctx context.Context, email string) (User, error){
  panic("not implemented")
}

func(s *service) Register(ctx context.Context, dto RegisterDTO)(User, error){
  panic("not implemented")
}

func(s *service) Login(ctx context.Context, dto LoginDTO)(User, error){
  user := User{
    Email: dto.Email,
    Password: HashPassword(dto.Password)
  }
  if err:= s.repo.Find(ctx, &user);err!=nil{
    return err
  }
  claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Minute * 2).Unix(),
	}
	token, err := GenerateToken(claims)
  return user, nil
}
