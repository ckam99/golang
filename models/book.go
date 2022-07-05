package models

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name      string         `json:"id" gorm:"type:varchar(50);NOT NULL"`
	Phone     string         `json:"phone" gorm:"type:varchar(50);NOT NULL;UNIQUE;UNIQUE_INDEX"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Book struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title"`
	Isbn      string         `json:"isbn" gorm:"index: , unique"`
	AuthorID  uint           `json:"author_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Author    Author         `json:"author"`
}
