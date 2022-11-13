package main

import (
	"fmt"
  "main/internal/domain/auth"
)

type User struct {
  ID int64
  Role auth.Role
}

func main() {
  user := User{
    Role: auth.Role(2),
  }
	fmt.Println(user)
}
