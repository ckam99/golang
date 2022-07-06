package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID               uint   `json:"id" gorm:"primaryKey; autoIncrement:true"`
	Name             string `faker:"name" json:"name" gorm:"type:varchar(60)"`
	Email            string `faker:"email" json:"email" gorm:"type:varchar(255);unique; NOT NULL"`
	Password         string
	Phone            string         `faker:"phone_number" json:"phone" gorm:"type:varchar(60);unique;NULL"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	EmailConfirmedAt time.Time      `json:"email_confirmed_at,omitempty" gorm:"autoCreateTime:false"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at"`
}

type UserSchema struct {
	ID               uint           `json:"id" `
	Name             string         `json:"name" `
	Email            string         `json:"email" `
	Phone            string         `json:"phone"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	EmailConfirmedAt time.Time      `json:"email_confirmed_at,omitempty" `
	DeletedAt        gorm.DeletedAt `json:"deleted_at"`
}

type UserUpdateSchema struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func ToUserReponse(u *User) *UserSchema {
	return &UserSchema{
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

func ToUserListResponse(users *[]User) *[]UserSchema {
	var newList []UserSchema
	for _, user := range *users {
		newList = append(newList, *ToUserReponse(&user))
	}
	return &newList
}
