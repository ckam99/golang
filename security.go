package main

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const secretKey = "ec52b01ce38571b4c4f0ecf157e52975f8f22bf580468aaddac98e0800a72d19"

func GenerateExampleToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["role"] = "admin"
	claims["email"] = "test@example.com"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	signedString, err := token.SignedString([]byte(secretKey))
	return signedString, err
}

func CreateAccessToken(user *User) (string, error) {
	claims := jwt.MapClaims{
		"email": user.Email,
		"role":  user.Role.Name,
		"exp":   time.Now().Add(time.Minute * 2).Unix(),
	}
	return GenerateJWT(claims)
}

func GetUserFromClaim(claims jwt.MapClaims) *User {
	user := User{
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

func GenerateJWT(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func VerifyJWT(decodedToken string) (*jwt.Token, error) {
	tokenString := DecodeJWT(decodedToken)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func DecodeJWT(bearToken string) string {
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}
	return ""
}
