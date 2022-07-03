package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email" gorm:"index:user_email_uniq , unique"`
}

type UserSchema struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserSchema struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (user *User) Serialize() UserSchema {
	return UserSchema{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
