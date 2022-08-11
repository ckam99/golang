package repository

import (
	"github.com/ckam225/golang/fiber/internal/entity"
	"github.com/ckam225/golang/fiber/internal/http/request"
	"github.com/ckam225/golang/fiber/internal/http/response"
	"github.com/ckam225/golang/fiber/internal/service"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FetchAllUsers() (*[]response.UserResponse, error) {
	var users []response.UserResponse
	if err := r.db.Model(&entity.User{}).Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *userRepository) GetAllUsers(p request.UserFilterParam) (*[]entity.User, error) {
	var users []entity.User
	if err := r.db.Offset(p.Skip).Limit(p.Limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *userRepository) CreateUser(obj *request.CreateUser) (*entity.User, error) {
	user := entity.User{
		Name:  obj.Name,
		Email: obj.Email,
		Phone: obj.Phone,
	}
	_, err := service.CreateUser(r.db, &user)
	return &user, err
}

func (r *userRepository) GetUser(user *entity.User) (*entity.User, error) {
	if err := r.db.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUserByID(id int) (*entity.User, error) {
	return service.GetUserByID(r.db, id)
}

func (r *userRepository) GetUserByEmail(email string) (*entity.User, error) {
	return service.GetUserByEmail(r.db, email)
}

func (r *userRepository) UpdateUser(userId int, payload *request.UpdateUser) (*entity.User, error) {
	if user, err := r.GetUserByID(userId); err != nil {
		return nil, err
	} else {
		err := r.db.Model(&user).Updates(entity.User{
			Name:  payload.Name,
			Phone: payload.Phone,
		}).Error
		return user, err
	}
}

func (r *userRepository) DeleteUser(userId uint, isSoftDelete bool) error {
	if !isSoftDelete {
		return r.db.Unscoped().Where("id = ?", userId).Delete(&entity.User{}).Error
	}
	return r.db.Where("id = ?", userId).Delete(&entity.User{}).Error
}

func (r *userRepository) CreateFakeUsers(maxLines int) error {
	return service.CreateFakeUsers(r.db, maxLines)
}

func (r *userRepository) CreateFakeUser() error {
	return service.CreateFakeUser(r.db)
}
