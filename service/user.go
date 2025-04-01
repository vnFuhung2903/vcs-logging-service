package service

import (
	"errors"
	"net/mail"

	"github.com/vnFuhung2903/vcs-logging-service/model"
	"github.com/vnFuhung2903/vcs-logging-service/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(email string, password string, name string) (*model.User, error)
	GetUserByEmail(email string) ([]*model.User, error)
}

type userService struct {
	Ur repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) UserService {
	return &userService{Ur: *ur}
}

func (userService *userService) Register(email string, password string, name string) (*model.User, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return nil, err
	}

	users, err := userService.Ur.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if len(users) > 0 {
		return nil, errors.New("email address is in used")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	user, err := userService.Ur.CreateUser(email, name, string(hash))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (userService *userService) GetUserByEmail(email string) ([]*model.User, error) {
	wallets, err := userService.Ur.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return wallets, nil
}
