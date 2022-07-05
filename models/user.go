package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"id"`
	Name       string      `json:"name"`
	Email      string      `json:"email" gorm:"index:user_email_uniq,unique"`
	Languages  []*Language `json:"languages" gorm:"many2many:user_languanges;"`
}

type Language struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"title"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Users     []*User        `json:"users" gorm:"many2many:user_languanges;"`
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
