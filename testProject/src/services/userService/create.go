package userService

import (
	"errors"
	"testProject/src/middleware"
	"testProject/src/models"
)

func (r *ServiceUser) CreateUser(email string, password string, name string, surName string) (*models.User, error) {
	user, err := r.repo.GetUserByEmail(email)

	if user != nil {
		return nil, errors.New("userController already exists")
	}

	salt, err := middleware.GenerateSalt()

	if err != nil {
		panic(err)
	}

	hashedPassword, err := middleware.HashPassword(password, salt)

	if err != nil {
		return nil, err
	}

	user = &models.User{
		Name:     name,
		Surname:  surName,
		Email:    email,
		Salt:     salt,
		Password: hashedPassword,
	}

	return r.repo.CreateUser(user)
}
