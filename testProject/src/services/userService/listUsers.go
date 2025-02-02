package userService

import (
	"testProject/src/models"
)

func (r *ServiceUser) ListUsers() (*[]models.User, error) {
	return r.repo.ListUsers()
}
