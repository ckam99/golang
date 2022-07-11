package security

import (
	"example/fiber/entity"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func CreateAccessToken(user *entity.User) (string, error) {
	claims := jwt.MapClaims{
		"email": user.Email,
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

func GenerateJWT(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func VerifyJWT(decodedToken string) (*jwt.Token, error) {
	tokenString := DecodeJWT(decodedToken)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
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
