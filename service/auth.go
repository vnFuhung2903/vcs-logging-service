package service

import (
	"errors"

	"github.com/vnFuhung2903/vcs-logging-service/model"
	"github.com/vnFuhung2903/vcs-logging-service/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(email string, password string) (*model.User, error)
}

type authService struct {
	Ur repository.UserRepository
}

func NewAuthService(ur *repository.UserRepository) AuthService {
	return &authService{Ur: *ur}
}

func (authService *authService) Login(email string, password string) (*model.User, error) {
	users, err := authService.Ur.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if len(users) > 1 {
		return nil, errors.New("more than one user found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return users[0], nil
}
