package repository

import (
	"example/fiber/entity"
	"example/fiber/http/request"
	"example/fiber/http/response"
	"example/fiber/service"

	"gorm.io/gorm"
)

type UserRepository struct {
	Query *gorm.DB
}

func (r *UserRepository) FetchAllUsers() (*[]response.UserResonse, error) {
	var users []response.UserResonse
	if err := r.Query.Model(&entity.User{}).Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *UserRepository) GetAllUsers(p request.UserFilterParam) (*[]entity.User, error) {
	var users []entity.User
	if err := r.Query.Offset(p.Skip).Limit(p.Limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *UserRepository) CreateUser(obj *request.CreateUser) (*entity.User, error) {
	user := entity.User{
		Name:  obj.Name,
		Email: obj.Email,
		Phone: obj.Phone,
	}
	_, err := service.CreateUser(r.Query, &user)
	return &user, err
}

func (r *UserRepository) GetUser(user *entity.User) (*entity.User, error) {
	if err := r.Query.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByID(id int) (*entity.User, error) {
	return service.GetUserByID(r.Query, id)
}

func (r *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	return service.GetUserByEmail(r.Query, email)
}

func (r *UserRepository) UpdateUser(userId int, payload *request.UpdateUser) (*entity.User, error) {
	if user, err := r.GetUserByID(userId); err != nil {
		return nil, err
	} else {
		err := r.Query.Model(&user).Updates(entity.User{
			Name:  payload.Name,
			Phone: payload.Phone,
		}).Error
		return user, err
	}
}

func (r *UserRepository) DeleteUser(userId uint, isSoftDelete bool) error {
	if !isSoftDelete {
		return r.Query.Unscoped().Where("id = ?", userId).Delete(&entity.User{}).Error
	}
	return r.Query.Where("id = ?", userId).Delete(&entity.User{}).Error
}

func (r *UserRepository) CreateFakeUsers(maxLines int) error {
	return service.CreateFakeUsers(r.Query, maxLines)
}

func (r *UserRepository) CreateFakeUser() error {
	return service.CreateFakeUser(r.Query)
}
