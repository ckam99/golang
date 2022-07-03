package models

import "time"

type Author struct {
	Name  string `gorm:"type:varchar(50);NOT NULL"`
	Phone string `gorm:"type:varchar(50);NOT NULL;UNIQUE;UNIQUE_INDEX"`
}

type Book struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Isbn      string `gorm:"index: , unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
