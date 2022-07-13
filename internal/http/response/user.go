package response

import (
	"project-struct/internal/entity"
	"time"

	"gorm.io/gorm"
)

// https://pkg.go.dev/github.com/go-playground/validator

type UserResonse struct {
	ID               uint           `json:"id" `
	Name             string         `json:"name" `
	Email            string         `json:"email" `
	Phone            string         `json:"phone"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	EmailConfirmedAt time.Time      `json:"email_confirmed_at,omitempty" `
	DeletedAt        gorm.DeletedAt `json:"deleted_at"`
}

func ParseUserEntity(u *entity.User) *UserResonse {
	return &UserResonse{
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

func ParseUserListEntity(users *[]entity.User) *[]UserResonse {
	var newList []UserResonse
	for _, user := range *users {
		newList = append(newList, *ParseUserEntity(&user))
	}
	return &newList
}
