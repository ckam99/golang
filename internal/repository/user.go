package repository

import (
	"github.com/ckam225/golang/fiber/internal/entity"
	"github.com/ckam225/golang/fiber/internal/security"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetAllUsers(limit, offset int) (*[]entity.User, error)
	CreateUser(obj *entity.User) (*entity.User, error)
	GetUser(user *entity.User) (*entity.User, error)
	GetUserByID(id int) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	GetUserByUniqueString(uniquId string) (*entity.User, error)
	UpdateUser(user *entity.User) (*entity.User, error)
	DeleteUser(userId uint, isSoftDelete bool) error
}

type userRepository struct {
	db *gorm.DB
}

func UserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetAllUsers(limit, offset int) (*[]entity.User, error) {
	var users []entity.User
	if err := r.db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *userRepository) CreateUser(user *entity.User) (*entity.User, error) {
	var err error
	user.Password, err = security.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	if err = r.db.Omit("email_confirmed_at").Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUser(user *entity.User) (*entity.User, error) {
	if err := r.db.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUserByID(id int) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByUniqueString(uniquId string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", uniquId).Or("phone = ?", uniquId).First(&user).Error
	return &user, err
}

func (r *userRepository) UpdateUser(payload *entity.User) (*entity.User, error) {
	if user, err := r.GetUserByID(int(payload.ID)); err != nil {
		return nil, err
	} else {
		err := r.db.Model(&user).Updates(payload).Error
		return user, err
	}
}

func (r *userRepository) DeleteUser(userId uint, isSoftDelete bool) error {
	if !isSoftDelete {
		return r.db.Unscoped().Where("id = ?", userId).Delete(&entity.User{}).Error
	}
	return r.db.Where("id = ?", userId).Delete(&entity.User{}).Error
}
