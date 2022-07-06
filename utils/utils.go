package utils

import (
	"encoding/json"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Print2Json(v any) {
	if s, e := json.MarshalIndent(v, "", "  "); e != nil {
		panic(e.Error())
	} else {
		fmt.Println(string(s))
	}
}

func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		panic(err.Error())
	}
	return string(hash)
}
