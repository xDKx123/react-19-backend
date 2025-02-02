package userRepository

import "testProject/src/models"

func (r *UserRepository) ListUsers() (*[]models.User, error) {
	// list users
	var users []models.User

	//user gorm to get all users
	r.db.Find(&users)

	//return users
	return &users, nil
}
