package service

import (
	"errors"
	"net/mail"

	"github.com/vnFuhung2903/postgresql/api"
	"github.com/vnFuhung2903/postgresql/model"
	"github.com/vnFuhung2903/postgresql/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(req *api.LoginReqBody) (*model.User, error)
}

type authService struct {
	Ur repository.UserRepository
}

func NewAuthService(ur *repository.UserRepository) AuthService {
	return &authService{Ur: *ur}
}

func (authService *authService) Login(req *api.LoginReqBody) (*model.User, error) {
	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		return nil, err
	}

	users, err := authService.Ur.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if len(users) > 1 {
		return nil, errors.New("more than one user found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	return users[0], nil
}
