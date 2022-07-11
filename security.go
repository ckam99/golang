package main

import (
	"errors"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type TokenMetadata struct {
	Expires int64
}

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
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}
	return GenerateJWT(claims)
}

func ExtractTokenUser(ctx *fiber.Ctx) (*User, error) {
	token := ctx.Locals("x-fiber-user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	if !token.Valid {
		return nil, errors.New("Token is invalid")
	}
	expires := int64(claims["exp"].(float64))
	println("token expire", expires)
	user := User{
		Email: claims["email"].(string),
	}
	//
	return &user, nil
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

func ExtractJWTMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	// Normally Authorization HTTP header.
	token, err := VerifyJWT(c.Get("Authorization"))
	if err != nil {
		return nil, err
	}
	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// Expires time.
		expires := int64(claims["exp"].(float64))
		return &TokenMetadata{
			Expires: expires,
		}, nil
	}
	return nil, err
}
