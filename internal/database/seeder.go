package database

import (
	"project-struct/internal/entity"

	"gorm.io/gorm"
)

func CreateRolesSeeder(db *gorm.DB) error {
	roles := []entity.Role{
		{Name: "superadmin"},
		{Name: "admin"},
		{Name: "user"},
	}
	err := db.Create(&roles).Error
	return err
}
