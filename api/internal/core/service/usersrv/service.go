package usersrv

import (
	"errors"

	"github.com/na7r1x/r8r/api/internal/core/domain"
	"github.com/na7r1x/r8r/api/internal/core/ports"
)

type service struct {
	userRepository ports.UserRepository
}

func New(userRepository ports.UserRepository) *service {
	return &service{
		userRepository: userRepository,
	}
}

func (srv *service) Register(user domain.User) error {
	return srv.userRepository.Set(user)
}

func (srv *service) Login(username string, password string) error {
	user, err := srv.userRepository.One(username)
	if err != nil {
		return errors.New("could not retrieve user [" + username + "]; reason: " + err.Error())
	}
	if user.Password == password {
		return nil
	} else {
		return errors.New("invalid password")
	}
}
