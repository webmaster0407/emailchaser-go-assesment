package main

import (
	"github.com/gin-gonic/gin"
	
	"emailchaser.com/backend-go/initializers"
	"emailchaser.com/backend-go/routes"
)

func main() {
	// Initialize environment variables and connect to the database
	if err := initializers.InitApp(); err != nil {
		panic(err)
	}

	// Create a new Gin engine with required middleware
	app := gin.Default()

	// Set up routes with the Gin engine
	routes.SetupRoutes(app)

	// Start the server
	if err := app.Run(":8080"); err != nil {
		panic(err)
	}
}