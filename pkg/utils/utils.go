package utils

import (
	"encoding/json"
	"errors"
	"log"
)

var ErrNoEntity = errors.New("entity not found")
var ErrUniqueField = errors.New("unique violation")
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrInvalidForeinKey = errors.New("foreign key violation")

func JSON(t any) {
	b, _ := json.MarshalIndent(t, "", " ")
	log.Println(string(b))
}
