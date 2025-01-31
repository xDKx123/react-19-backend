package userController

import (
	"github.com/gin-gonic/gin"
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
}

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (c *UserController) Register(ctx *gin.Context) {
	var reqParams CreateUserRequest

	if err := ctx.ShouldBindJSON(&reqParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.service.CreateUser(reqParams.Email, reqParams.Password, reqParams.Name, reqParams.Surname)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
