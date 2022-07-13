package service

import (
	"example/fiber/entity"

	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
)

func CreateFakeUsers(db *gorm.DB, maxLines int) error {
	var err error
	for i := 0; i < maxLines; i++ {
		//if e := db.Omit("email_confirmed_at", "deleted_at", "id").Create(&entity.User{}).Error; e != nil {
		if e := CreateFakeUser(db); e != nil {
			err = e
			break
		}
	}
	return err
}

func CreateFakeUser(db *gorm.DB) error {
	if err := db.Omit("email_confirmed_at").Create(&entity.User{
		Name:  faker.Name(),
		Email: faker.Email(),
		Phone: faker.Phonenumber(),
	}).Error; err != nil {
		return err
	}
	return nil
}
