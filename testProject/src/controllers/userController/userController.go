package userController

import (
	"net/http"
	"testProject/src/services/userServices"
)

type UserController struct {
	service *userServices.UserService
}

func NewUserController(service *userServices.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var 
}