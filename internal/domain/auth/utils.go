package auth

import (
  "github.com/golang-jwt/jwt/v4"
"golang.org/x/crypto/bcrypt"
)

func GenerateToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func GetUserFromClaim(claims jwt.MapClaims) *User {
	user := entity.User{
		Email: claims["email"].(string),
	}
	return &user
}

// extract jwt
func ExtractToken(ctx *fiber.Ctx) (*jwt.MapClaims, error) {
	// Normally Authorization HTTP header.
	token, err := VerifyToken(c.Get("Authorization"))
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	return &claims, err
}

// verify token
func VerifyToken(encodedToken string) (*jwt.Token, error) {
	tokenString, err := DecodeToken(encodedToken)
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// decode token
func DecodeToken(bearToken string) (string, error) {
	tokenArray := strings.Split(bearToken, " ")
	accessToken := ""
	if len(tokenArray) == 2 {
		accessToken = tokenArray[1]
	}
	if accessToken == "" {
		return "", errors.New("token is empty or invalid")
	}
	return accessToken, nil
}


// get auth user
func GetAuthUser(db *gorm.DB, ctx *fiber.Ctx) (*entity.User, error) {
	claims, err := ExtractJWT(ctx)
	//claims, err := ExtractJsonWebToken(ctx)
	if err != nil {
		return nil, err
	}
	user := GetUserFromClaim(*claims)
	if err = db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// hash password 

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// verify password
func VerifyPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return !(err != nil)
}