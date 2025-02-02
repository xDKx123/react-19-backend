package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"testProject/src/controllers"
	"testProject/src/database"
	"testProject/src/middleware"
	"testProject/src/migrations"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	log.Println("Starting server...")
	db, err := database.ConnectLocalDatabase()

	if err != nil || db == nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	err = migrations.MigrateModels(db)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	fmt.Println("Successfully migrated models!")

	// Create a new Gin router
	fmt.Println("Initializing Gin router...")
	router := gin.Default()

	// Apply the middleware to the router
	router.Use(middleware.JSONResponseGinMiddleware())

	//Initializing static routes
	router.Static("/static", "./static")

	//TODO: move to a utility controller
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to the Home Page! Access static files at /static/")
	})

	//Ping page
	//TODO: move to a utility controller
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	controllers.InitializeControllers(db, router)

	routerErr := router.Run(":8080")

	if routerErr != nil {
		log.Fatalf("Failed to start server: %v", routerErr)
	}

	fmt.Println("Server is running on port 8080")
}
