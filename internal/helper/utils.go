package helper

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func PrintJson(v any) {
	if s, e := json.MarshalIndent(v, "", "  "); e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println(string(s))
	}
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func SetHttpError(msg string) *fiber.Map {
	return &fiber.Map{
		"message": msg,
	}
}
