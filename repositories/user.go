package repositories

import (
	"errors"
	"example/gorm/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func (repo *UserRepository) GetUsers(limit, offset int) (*[]models.User, error) {
	users := &[]models.User{}
	if err := repo.Db.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) GetUser(user *models.User) (*models.User, error) {
	if err := repo.Db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetUserByID(id int) (*models.User, error) {
	user := &models.User{}
	if err := repo.Db.Find(user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return user, errors.New("user does'nt exists")
	}
	return user, nil
}

func (repo *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	if err := repo.Db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	if err := repo.Db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) CreateUsers(users *[]models.User) (*[]models.User, error) {
	err := repo.Db.Create(&users).Error
	return users, err
}

func (repo *UserRepository) UpdateUser(id int, payload models.UpdateUserSchema) (*models.User, error) {
	user, err := repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	user.Name = payload.Name
	repo.Db.Save(&user)

	return user, err
}

// Retrieve user list with edger loading languages
func (repo *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := repo.Db.Model(&models.User{}).Preload("Languages").Find(&users).Error
	return users, err
}

// Retrieve language list with edger loading users
func (repo *UserRepository) GetAllLanguages() ([]models.Language, error) {
	var languages []models.Language
	err := repo.Db.Model(&models.Language{}).Preload("Users").Find(&languages).Error
	return languages, err
}
