package main

import (
	"fmt"
	"net/http"

	"example.com/rest-api/db"
	"example.com/rest-api/models"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize the database
	db.InitDB()

	// Create a new server
	server := gin.Default()

	// Define a GET request route
	server.GET("/events", getEvents)

	// Define a POST request route to add an event
	server.POST("/events", addEvent)
	// Run the server on port 8080
	server.Run(":8080")



	fmt.Println("hello from vim")
}


// Function that returns events
func getEvents(c *gin.Context) {
	// Create a slice of events
	events := models.GetAllEvents()

	// Return the events as JSON
	c.JSON(http.StatusOK, events)
}


// Function that adds an event
func addEvent(c *gin.Context) {
	// Create a new event
	var event models.Event

	// Bind the request body to the event
	if err := c.ShouldBindJSON(&event); err != nil {
		// Return an error if the event cannot be bound
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ID = fmt.Sprintf("%d", len(models.GetAllEvents())+1)
	event.UserID = "1"

	// Save the event
	event.Save()

	// Return the event
	c.JSON(http.StatusCreated, gin.H{ 
		"message": "Your event was succesfully created!" , 
		"event": event,
	})
}
