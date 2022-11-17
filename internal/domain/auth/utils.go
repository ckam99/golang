package auth

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var TOKEN_EXPIRE_TIME = time.Now().Add(time.Hour * 1).Unix()

// Generate new token
func GenerateToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

// verify token
func VerifyToken(encodedToken string) (*jwt.Token, error) {
	// token := ctx.Locals("x-fiber-user").(*jwt.Token)
	// claims := token.Claims.(jwt.MapClaims)
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
