package service

import (
	"github.com/ckam225/golang/fiber/internal/entity"
	"github.com/ckam225/golang/fiber/internal/security"
	"github.com/ckam225/golang/fiber/internal/utils"
	"github.com/ckam225/golang/fiber/pkg/mailer"

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

func SendConfirmationEmail(db *gorm.DB, user *entity.User) error {
	vcode := entity.Verycode{
		Email: user.Email,
		Code:  utils.GenerateHashCode(),
	}
	if err := db.Create(&vcode).Error; err != nil {
		return err
	}
	mailer.NotificationChannel <- &mailer.Notification{
		To: []string{user.Email},
		Data: map[string]string{
			"email": user.Email,
			"name":  user.Name,
			"code":  vcode.Code,
		},
		Subject:  "Registration",
		Template: "mail/confirm_action.tmpl",
	}
	return nil
}
