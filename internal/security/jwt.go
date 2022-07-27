package security

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/ckam225/golang/fiber/internal/entity"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func CreateAccessToken(user *entity.User) (string, error) {
	claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Minute * 2).Unix(),
	}
	return GenerateJWT(claims)
}

func RefreshAccessToken(user *entity.User, oldAccessToken string) (string, error) {
	claims := jwt.MapClaims{
		"email": user.Email,
		"token": oldAccessToken,
		"exp":   time.Now().Add(time.Minute * 2).Unix(),
	}
	return GenerateJWT(claims)
}

func GetUserFromClaim(claims jwt.MapClaims) *entity.User {
	user := entity.User{
		Email: claims["email"].(string),
	}
	return &user
}

func ExtractJWT(ctx *fiber.Ctx) (*jwt.MapClaims, error) {
	token := ctx.Locals("x-fiber-user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	// expires := int64(claims["exp"].(float64))
	return &claims, nil
}

func ExtractJsonWebToken(c *fiber.Ctx) (*jwt.MapClaims, error) {
	// Normally Authorization HTTP header.
	token, err := VerifyJWT(c.Get("Authorization"))
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	return &claims, err
}

func VerifyJWT(decodedToken string) (*jwt.Token, error) {
	tokenString, err := DecodeJWT(decodedToken)
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

func DecodeJWT(bearToken string) (string, error) {
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
