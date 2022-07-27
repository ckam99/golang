package security

import (
	"github.com/ckam225/golang/fiber/internal/entity"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

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

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func VerifyPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return !(err != nil)
}
