package repository

import (
	"example/fiber/entity"
	"example/fiber/utils"

	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
)

type UserRepository struct {
	Query  *gorm.DB
	Filter UserFilterParam
}

type UserFilterParam struct {
	Limit int
	Skip  int
}

func (r *UserRepository) FetchAllUsers() (*[]entity.UserSchema, error) {
	var users []entity.UserSchema
	if err := r.Query.Model(&entity.User{}).Limit(r.Filter.Limit).Offset(r.Filter.Skip).Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *UserRepository) GetAllUsers() (*[]entity.User, error) {
	var users []entity.User
	if err := r.Query.Limit(r.Filter.Limit).Offset(r.Filter.Skip).Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *UserRepository) CreateUser(user *entity.User) (*entity.User, error) {
	if user.Password != "" {
		user.Password = utils.HashPassword(user.Password)
	}
	if err := r.Query.Omit("email_confirmed_at").Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUser(user *entity.User) (*entity.User, error) {
	if err := r.Query.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByID(id uint) (*entity.User, error) {
	var user entity.User
	if err := r.Query.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.Query.Where("email = ?", email).First(&user).Error; err != nil {
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
		err := r.Query.Save(&user).Error
		return user, err
	}
}

func (r *UserRepository) DeleteUser(userId uint, permanlty bool) error {
	if permanlty {
		return r.Query.Unscoped().Where("id = ?", userId).Delete(&entity.User{}).Error
	}
	return r.Query.Where("id = ?", userId).Delete(&entity.User{}).Error
}

func (r *UserRepository) ChangeUserPassword(user *entity.User, newPassword string) error {
	user.Password = utils.HashPassword(newPassword)
	err := r.Query.Save(&user).Error
	return err
}

func (r *UserRepository) CreateFakeUsers(maxLines int) error {
	var err error
	for i := 0; i < maxLines; i++ {
		if e := r.Query.Omit("email_confirmed_at", "deleted_at", "id").Create(&entity.User{}).Error; e != nil {
			err = e
			break
		}
	}
	return err
}

func (r *UserRepository) CreateFakeUser() error {
	if err := r.Query.Omit("email_confirmed_at").Create(&entity.User{
		Name:  faker.Name(),
		Email: faker.Email(),
		Phone: faker.Phonenumber(),
	}).Error; err != nil {
		return err
	}
	return nil
}
