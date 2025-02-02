package userController

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testProject/src/middleware"
	"time"
)

type UserLoginRequest struct {
	UsernameOrEmail string `json:"usernameOrEmail" binding:"required"`
	Password        string `json:"password" binding:"required"`
}

func (c *UserController) Login(ctx *gin.Context) {
	var reqParams UserLoginRequest

	if err := ctx.ShouldBindJSON(&reqParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.service.Login(reqParams.UsernameOrEmail, reqParams.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := middleware.GenerateToken(user.ID, user.Email, time.Hour)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
