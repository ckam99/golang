package models

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name      string `gorm:"type:varchar(50);NOT NULL"`
	Phone     string `gorm:"type:varchar(50);NOT NULL;UNIQUE;UNIQUE_INDEX"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Book struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Isbn      string `gorm:"index: , unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
