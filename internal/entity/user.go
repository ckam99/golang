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
	Roles            []*Role        `json:"roles" gorm:"many2many:user_roles"`
}

type Role struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(60)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Verycode struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"type:varchar(255)"`
	Target    string    `json:"target" gorm:"type:varchar(60)"`
	Code      string    `json:"code" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) HasRole(role string) bool {
	for _, r := range u.Roles {
		if r.Name == role {
			return true
		}
	}
	return false
}
