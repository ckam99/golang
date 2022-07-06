package schema

import (
	"example/fiber/entity"
	"time"

	"gorm.io/gorm"
)

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
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
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
