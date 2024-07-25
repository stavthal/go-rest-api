package main

import (
	"fmt"
	"os"

	"example.com/rest-api/db"
	"example.com/rest-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load(".env.local")

	// Get mode from .env file
	gin_mode := os.Getenv("GIN_MODE")

	// Set the GIN_MODE to release/debug
	gin.SetMode(gin_mode)

	if err != nil {
		panic("Error loading .env file")
	}

	// Initialize the database
	db.InitDB()

	// Create a new server
	server := gin.Default()



	// Register the routes
	routes.RegisterRoutes(server)

	// Run the server on port 8080
	server.Run(":8080")



	fmt.Println("hello from vim")
}


