package service

import (
	"errors"
	"net/mail"

	"github.com/vnFuhung2903/postgresql/api"
	"github.com/vnFuhung2903/postgresql/model"
	"github.com/vnFuhung2903/postgresql/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(req *api.RegisterReqBody) (*model.User, error)
}

type userService struct {
	Ur repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) UserService {
	return &userService{Ur: *ur}
}

func (userService *userService) Register(req *api.RegisterReqBody) (*model.User, error) {
	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		return nil, err
	}

	users, err := userService.Ur.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if len(users) > 0 {
		return nil, errors.New("email address is in used")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return nil, err
	}

	user, err := userService.Ur.CreateUser(req.Email, req.Name, string(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}
