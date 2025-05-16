package repositories

import (
	"net/mail"

	"github.com/vnFuhung2903/vcs-logging-services/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	CreateUser(email string, password string) (*models.User, error)
	UpdateEmail(user *models.User, email string) error
	UpdatePassword(user *models.User, password string) error
	DeleteUser(email string) error
}

type userRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{Db: db}
}

func (ur *userRepository) FindById(id uint) (*models.User, error) {
	var user models.User
	res := ur.Db.First(&user, models.User{Id: id})
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (ur *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	res := ur.Db.First(&user, models.User{Email: email})
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (ur *userRepository) CreateUser(email string, password string) (*models.User, error) {
	newUser := &models.User{
		Email:    email,
		Password: password,
	}
	res := ur.Db.Create(newUser)
	if res.Error != nil {
		return nil, res.Error
	}
	return newUser, nil
}

func (ur *userRepository) UpdateEmail(user *models.User, email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return err
	}
	res := ur.Db.Model(user).Update("email", email)
	return res.Error
}

func (ur *userRepository) UpdatePassword(user *models.User, password string) error {
	res := ur.Db.Model(user).Update("password", password)
	return res.Error
}

func (ur *userRepository) DeleteUser(email string) error {
	res := ur.Db.Where("email = ?", email).Delete(&models.User{})
	return res.Error
}
