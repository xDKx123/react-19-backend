package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"testProject/src/controllers/userController"
)

func InitializeControllers(db *gorm.DB, router *gin.Engine) {
	userController.InitializeUserControllers(db, router)
}
