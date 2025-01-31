package userServices

import (
	"errors"
	"testProject/src/middleware"
	"testProject/src/models"
	userRepository "testProject/src/repository/user"
)

type UserService struct {
	repo *userRepository.UserRepository
}

func NewUserService(repo *userRepository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (r *UserService) Login(email string, password string) (*models.User, error) {
	user, err := r.repo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("Invalid username or password")
	}

	if !user.CheckPassword(password) {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (r *UserService) CreateUser(email string, password string, name string, surName string) (*models.User, error) {
	user, err := r.repo.GetUserByEmail(email)

	if user != nil {
		return nil, errors.New("User already exists")
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
