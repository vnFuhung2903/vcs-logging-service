package repository

import (
	"net/mail"

	"github.com/vnFuhung2903/vcs-logging-service/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindById(id uint) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	CreateUser(email string, password string) (*model.User, error)
	UpdateEmail(user *model.User, email string) error
	UpdatePassword(user *model.User, password string) error
	DeleteUser(email string) error
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

func (ur *userRepository) FindById(id uint) (*model.User, error) {
	var user model.User
	res := ur.Db.First(&user, model.User{Id: id})
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (ur *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	res := ur.Db.First(&user, model.User{Email: email})
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (ur *userRepository) CreateUser(email string, password string) (*model.User, error) {
	newUser := &model.User{
		Email:    email,
		Password: password,
	}
	res := ur.Db.Create(newUser)
	if res.Error != nil {
		return nil, res.Error
	}
	return newUser, nil
}

func (ur *userRepository) UpdateEmail(user *model.User, email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return err
	}
	res := ur.Db.Model(user).Update("email", email)
	return res.Error
}

func (ur *userRepository) UpdatePassword(user *model.User, password string) error {
	res := ur.Db.Model(user).Update("password", password)
	return res.Error
}

func (ur *userRepository) DeleteUser(email string) error {
	res := ur.Db.Where("email = ?", email).Delete(&model.User{})
	return res.Error
}
