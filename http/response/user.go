package response

import (
	"example/fiber/entity"
	"time"
)

// https://pkg.go.dev/github.com/go-playground/validator

type UserResponse struct {
	ID               uint      `json:"id" `
	Name             string    `json:"name" `
	Email            string    `json:"email" `
	Phone            string    `json:"phone"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	EmailConfirmedAt time.Time `json:"email_confirmed_at,omitempty" `
	DeletedAt        time.Time `json:"deleted_at"`
}

type AccessToken struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}

func ParseUserEntity(u *entity.User) *UserResponse {
	return &UserResponse{
		ID:               u.ID,
		Name:             u.Name,
		Email:            u.Email,
		Phone:            u.Phone,
		CreatedAt:        u.CreatedAt,
		UpdatedAt:        u.UpdatedAt,
		EmailConfirmedAt: u.EmailConfirmedAt,
		DeletedAt:        u.DeletedAt.Time,
	}
}

func ParseUserListEntity(users *[]entity.User) *[]UserResponse {
	var newList []UserResponse
	for _, user := range *users {
		newList = append(newList, *ParseUserEntity(&user))
	}
	return &newList
}
