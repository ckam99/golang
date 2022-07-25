package database

import (
	"github.com/ckam225/golang/fiber/internal/entity"

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

func DatabaseSeeder(db *gorm.DB) error {
	if err := CreateRolesSeeder(db); err != nil {
		return err
	}
	return nil
}
