package auth

import (
	"time"
)

type User struct {
	ID                int64   `json:"id"`
	FullName          string  `json:"full_name"`
	Email             string  `json:"email"`
	Phone             *string `json:"phone"`
	Password          *string
	Role              string     `json:"role"`
	IsActive          *bool      `json:"is_active"`
	EmailConfirmedAt  *time.Time `json:"email_confirm_at"`
	PhoneConfirmedAt  *time.Time `json:"phone_confirm_at"`
	PasswordChangedAt *time.Time `json:"password_changed_at"`
	CreatedAt         *time.Time `json:"created_at"`
	UpatedAt          *time.Time `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at"`
}
