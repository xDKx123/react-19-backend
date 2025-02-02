package userController

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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

	ctx.JSON(http.StatusOK, gin.H{"userController": user})
}
