package userController

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"testProject/src/middleware"
	"testProject/src/repository/userRepository"
	"testProject/src/services/userService"
)

type UserController struct {
	service *userService.ServiceUser
}

func NewUserController(service *userService.ServiceUser) *UserController {
	return &UserController{service: service}
}

func InitializeUserControllers(db *gorm.DB, router *gin.Engine) {
	userRepo := userRepository.NewUserRepository(db)
	userSrv := userService.NewUserService(userRepo)
	userCtrl := NewUserController(userSrv)

	router.POST("/userController/create", userCtrl.Register)
	router.POST("/auth/login", userCtrl.Login)
	//router.GET("/users/list", middleware.AuthMiddlewareV2(), userCtrl.List)

	protectedRoutes := router.Group("")
	protectedRoutes.Use(middleware.AuthMiddlewareV2())

	protectedRoutes.GET("/users/list", userCtrl.listUsers)

}
