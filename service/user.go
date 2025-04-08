package service

import (
	"net/mail"

	"github.com/vnFuhung2903/vcs-logging-service/model"
	"github.com/vnFuhung2903/vcs-logging-service/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(email string, password string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
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

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	user, err := userService.Ur.CreateUser(email, string(hash))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (userService *userService) GetUserByEmail(email string) (*model.User, error) {
	user, err := userService.Ur.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
