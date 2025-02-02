package userController

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *UserController) listUsers(ctx *gin.Context) {
	users, err := c.service.ListUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
	return
}
