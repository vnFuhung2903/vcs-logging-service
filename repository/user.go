package repository

import (
	"github.com/vnFuhung2903/vcs-logging-service/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindById(id uint) ([]*model.User, error)
	FindByEmail(email string) (*model.User, error)
	CreateUser(email string, password string) (*model.User, error)
	UpdateEmail(user *model.User, email string) error
	UpdateName(user *model.User, name string) error
	UpdatePassword(user *model.User, password string) error
}

type userRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{Db: db}
}

func (ur *userRepository) FindAll() ([]*model.User, error) {
	var users []*model.User
	res := ur.Db.Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}
	return users, nil
}

func (ur *userRepository) FindById(id uint) ([]*model.User, error) {
	var users []*model.User
	res := ur.Db.Find(&users, model.User{Id: id})
	if res.Error != nil {
		return nil, res.Error
	}
	return users, nil
}

func (ur *userRepository) FindByEmail(email string) (*model.User, error) {
	var users *model.User
	res := ur.Db.First(&users, model.User{Email: email})
	if res.Error != nil {
		return nil, res.Error
	}
	return users, nil
}

func (ur *userRepository) CreateUser(email string, password string) (*model.User, error) {
	res := ur.Db.Create(model.User{
		Email:    email,
		Password: password,
	})
	if res.Error != nil {
		return nil, res.Error
	}

	var user *model.User
	res = ur.Db.First(&user, model.User{Email: email})
	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (ur *userRepository) UpdateEmail(user *model.User, email string) error {
	res := ur.Db.Save(model.User{
		Id:       user.Id,
		Email:    email,
		Password: user.Password,
	})
	return res.Error
}

func (ur *userRepository) UpdateName(user *model.User, name string) error {
	res := ur.Db.Save(model.User{
		Id:       user.Id,
		Email:    user.Email,
		Password: user.Password,
	})
	return res.Error
}

func (ur *userRepository) UpdatePassword(user *model.User, password string) error {
	res := ur.Db.Save(model.User{
		Id:       user.Id,
		Email:    user.Email,
		Password: password,
	})
	return res.Error
}

func (ur *userRepository) DeleteUser(user *model.User) error {
	res := ur.Db.Delete(model.User{}, user.Id)
	return res.Error
}
