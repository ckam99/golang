package schema

import (
	"example/fiber/entity"
	"time"

	"gorm.io/gorm"
)

// https://pkg.go.dev/github.com/go-playground/validator

type User struct {
	ID               uint           `json:"id" `
	Name             string         `json:"name" `
	Email            string         `json:"email" `
	Phone            string         `json:"phone"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	EmailConfirmedAt time.Time      `json:"email_confirmed_at,omitempty" `
	DeletedAt        gorm.DeletedAt `json:"deleted_at"`
}

type UserUpdate struct {
	Name  string `validate:"required,min=2,max=60"`
	Phone string `validate:"min=6,max=60"`
}

type UserRegister struct {
	Name     string `validate:"required,min=2,max=60"`
	Phone    string `validate:"min=6,max=60"`
	Email    string `validate:"required,email,max=255"`
	Password string `validate:"required,min=6"`
}

type UserFilterParam struct {
	Limit int
	Skip  int
}

func UserReponse(u *entity.User) *User {
	return &User{
		ID:               u.ID,
		Name:             u.Name,
		Email:            u.Email,
		Phone:            u.Phone,
		CreatedAt:        u.CreatedAt,
		UpdatedAt:        u.UpdatedAt,
		EmailConfirmedAt: u.EmailConfirmedAt,
		DeletedAt:        u.DeletedAt,
	}
}

func UserListResponse(users *[]entity.User) *[]User {
	var newList []User
	for _, user := range *users {
		newList = append(newList, *UserReponse(&user))
	}
	return &newList
}
