package userService

import (
	"errors"
	"testProject/src/models"
)

func (r *ServiceUser) Login(email string, password string) (*models.User, error) {
	user, err := r.repo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	if !user.CheckPassword(password) {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
