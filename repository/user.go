package repository

import (
	"example/fiber/entity"
	"example/fiber/utils"

	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
)

type UserRepository struct {
	Q *gorm.DB
}

type UserQueryParam struct {
	Limit int
	Skip  int
}

func (r *UserRepository) GetAllUsers(q *UserQueryParam) (*[]entity.User, error) {
	var users []entity.User
	if err := r.Q.Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *UserRepository) CreateUser(user *entity.User) (*entity.User, error) {
	if user.Password != "" {
		user.EncryptPassword()
	}
	if err := r.Q.Omit("email_confirmed_at").Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUser(user *entity.User) (*entity.User, error) {
	if err := r.Q.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByID(id uint) (*entity.User, error) {
	var user entity.User
	if err := r.Q.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.Q.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(userId uint, payload *entity.UserUpdateSchema) (*entity.User, error) {
	if user, err := r.GetUserByID(userId); err != nil {
		return nil, err
	} else {
		user.Name = payload.Name
		user.Phone = payload.Phone
		err := r.Q.Save(&user).Error
		return user, err
	}
}

func (r *UserRepository) DeleteUser(userId uint, permanlty bool) error {
	if permanlty {
		return r.Q.Unscoped().Where("id = ?", userId).Delete(&entity.User{}).Error
	}
	return r.Q.Where("id = ?", userId).Delete(&entity.User{}).Error
}

func (r *UserRepository) ChangeUserPassword(user *entity.User, newPassword string) error {
	user.Password = utils.HashPassword(newPassword)
	err := r.Q.Save(&user).Error
	return err
}

func (r *UserRepository) CreateFakeUsers(maxLines int) error {
	var err error
	for i := 0; i < maxLines; i++ {
		if e := r.Q.Omit("email_confirmed_at", "deleted_at", "id").Create(&entity.User{}).Error; e != nil {
			err = e
			break
		}
	}
	return err
}

func (r *UserRepository) CreateFakeUser() error {
	if err := r.Q.Omit("email_confirmed_at").Create(&entity.User{
		Name:  faker.Name(),
		Email: faker.Email(),
		Phone: faker.Phonenumber(),
	}).Error; err != nil {
		return err
	}
	return nil
}
