package userServices

import (
	"errors"
	"testProject/src/models"
	userRepository "testProject/src/repository/user"
)

type UserService struct {
	repo *userRepository.UserRepository
}

func (r *UserService) Login(email, password string) (*models.User, error) {
	user, err := r.repo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("Invalid username or password")
	}

	if !user.CheckPassword(password) {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
