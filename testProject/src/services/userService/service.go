package userService

import (
	"testProject/src/repository/userRepository"
)

type ServiceUser struct {
	repo *userRepository.UserRepository
}

func NewUserService(repo *userRepository.UserRepository) *ServiceUser {
	return &ServiceUser{repo: repo}
}
