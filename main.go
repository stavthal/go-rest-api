package main

import (
	"fmt"

	"example.com/rest-api/db"
	"example.com/rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {

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


