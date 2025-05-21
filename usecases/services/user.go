package services

import (
	"errors"
	"net/mail"

	"github.com/vnFuhung2903/vcs-logging-services/models"
	"github.com/vnFuhung2903/vcs-logging-services/usecases/repositories"
)

type UserService interface {
	Register(email string, password string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Update(user *models.User, key string, newData string) error
	Delete(email string) error
}

type userService struct {
	Ur repositories.UserRepository
}

func NewUserService(ur repositories.UserRepository) UserService {
	return &userService{Ur: ur}
}

func (userService *userService) Register(email string, password string) (*models.User, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return nil, err
	}

	user, err := userService.Ur.CreateUser(email, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (userService *userService) FindByEmail(email string) (*models.User, error) {
	user, err := userService.Ur.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (userService *userService) Update(user *models.User, key string, newData string) error {
	var err error
	switch key {
	case "email":
		err = userService.Ur.UpdateEmail(user, newData)
	case "password":
		err = userService.Ur.UpdatePassword(user, newData)
	default:
		err = errors.New("invalid key")
	}
	return err
}

func (userService *userService) Delete(email string) error {
	err := userService.Ur.DeleteUser(email)
	return err
}
