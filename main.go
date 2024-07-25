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

	// GET request route to get a single event
	server.GET("/events/:id", getEvent)

	// Define a POST request route to add an event
	server.POST("/events", addEvent)





	// Run the server on port 8080
	server.Run(":8080")



	fmt.Println("hello from vim")
}


// Function that returns events
func getEvents(c *gin.Context) {
	// Create a slice of events
	events, err := models.GetAllEvents()

	if err != nil {
		// Return an error if the events cannot be retrieved
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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

	// Set the ID and UserID of the event
	events, _ := models.GetAllEvents()

	event.ID = fmt.Sprintf("%d", len(events)+1)
	event.UserID = "1"

	// Save the event
	err := event.Save()

	if err != nil {
		// Return an error if the event cannot be saved
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the event
	c.JSON(http.StatusCreated, gin.H{ 
		"message": "Your event was succesfully created!" , 
		"event": event,
	})
}

func getEvent(c *gin.Context) {
	// Get the ID from the URL
	id := c.Param("id")

	// Get the event from the database
	event, err := models.GetEventById(id)

	if err != nil {
		// Return an error if the event cannot be retrieved
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If Event is empty, return a 404
	if event.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	// Return the event as JSON
	c.JSON(http.StatusOK, event)

}
