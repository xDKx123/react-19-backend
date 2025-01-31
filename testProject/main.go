package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"testProject/src/controllers/userController"
	"testProject/src/database"
	"testProject/src/middleware"
	"testProject/src/models"
	"testProject/src/repository/user"
	"testProject/src/services/userServices"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	log.Println("Starting server...")
	db, err := database.ConnectLocalDatabase()
	// Run migrations
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	fmt.Println("Successfully migrated models!")

	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Example of a simple home page handler
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Welcome to the Home Page! Access static files at /static/"))
	//})

	//err = http.ListenAndServe(":8080", nil)

	//if err != nil {
	//	log.Fatalf("Failed to start server: %v", err)
	//}

	// Create a new Gin router

	fmt.Println("Initializing Gin router...")
	router := gin.Default()

	// Apply the middleware to the router
	router.Use(middleware.JSONResponseGinMiddleware())

	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to the Home Page! Access static files at /static/")
	})

	router.GET("/ping", func(c *gin.Context) {
		//json response
		c.JSON(200, gin.H{
			"message": "pong ",
		})
	})

	// Initialize controllers
	userRepo := user.NewUserRepository(db)
	userSrv := userServices.NewUserService(userRepo)
	userCtrl := userController.NewUserController(userSrv)
	router.POST("/user/create", userCtrl.Register)

	routerErr := router.Run(":8080")

	if routerErr != nil {
		log.Fatalf("Failed to start server: %v", routerErr)
	}

	fmt.Println("Server is running on port 8080")
}
