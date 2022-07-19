package service

import (
	"example/fiber/entity"
	"example/fiber/security"

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
	hash, err := security.HashPassword("password")
	if err != nil {
		return err
	}
	err = db.Omit("email_confirmed_at").Create(&entity.User{
		Name:     faker.Name(),
		Email:    faker.Email(),
		Phone:    faker.Phonenumber(),
		Password: hash,
	}).Error
	return err
}

func GetUserByID(db *gorm.DB, id int) (*entity.User, error) {
	var user entity.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(db *gorm.DB, email string) (*entity.User, error) {
	var user entity.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUniqueString(db *gorm.DB, uniquId string) (*entity.User, error) {
	var user entity.User
	err := db.Where("email = ?", uniquId).Or("phone = ?", uniquId).First(&user).Error
	return &user, err
}

func CreateUser(db *gorm.DB, user *entity.User) (*entity.User, error) {
	var err error
	user.Password, err = security.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	if err = db.Omit("email_confirmed_at").Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
