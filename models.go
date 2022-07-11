package main

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	AccessToken string `json:"access_token"`
	ExpireAt    string `json:"expire_at"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token Token  `json:"token"`
	Role  Role   `json:"role"`
}

type Role struct {
	Name string `json:"name"`
}
