package utils

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func PrintJson(v any) {
	if s, e := json.MarshalIndent(v, "", "  "); e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println(string(s))
	}
}

func SetHttpError(msg string) *fiber.Map {
	return &fiber.Map{
		"message": msg,
	}
}
