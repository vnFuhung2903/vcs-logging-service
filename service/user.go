package service

import (
	"errors"
	"net/mail"

	"github.com/vnFuhung2903/vcs-logging-service/model"
	"github.com/vnFuhung2903/vcs-logging-service/repository"
)

type UserService interface {
	Register(email string, password string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Update(user *model.User, key string, newData string) error
	Delete(user *model.User) error
}

type userService struct {
	Ur repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) UserService {
	return &userService{Ur: *ur}
}

func (userService *userService) Register(email string, password string) (*model.User, error) {
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

func (userService *userService) FindByEmail(email string) (*model.User, error) {
	user, err := userService.Ur.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (userService *userService) Update(user *model.User, key string, newData string) error {
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

func (userService *userService) Delete(user *model.User) error {
	err := userService.Ur.DeleteUser(user)
	return err
}
