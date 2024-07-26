package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

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

	event.ID = int64(len(events) + 1)

	// Get the user ID
	userId := context.GetInt64("userId")

	// Convert userId to string and set it as the UserID of the event
	event.UserID = userId

	// Save the event
	err = event.Save()

	if err != nil {
		// Return an error if the event cannot be saved
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the event
	c.JSON(http.StatusCreated, gin.H{
		"message": "Your event was succesfully created!",
		"event":   event,
	})
}

func getEvent(c *gin.Context) {
	// Get the ID from the URL
	id := c.Param("id")

	// Convert the ID to an integer
	convertedId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		// Return an error if the ID cannot be converted
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Get the event from the database
	event, err := models.GetEventById(convertedId)

	// If Event is empty, return a 404
	if event.ID == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if err != nil {
		// Return an error if the event cannot be retrieved
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the event as JSON
	c.JSON(http.StatusOK, event)

}

func updateEvent(c *gin.Context) {
	// Get the ID from the URL
	id := c.Param("id")

	// Convert the ID to an integer
	convertedId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		// Return an error if the ID cannot be converted
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	event, err := models.GetEventById(convertedId)

	if event.ID == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if err != nil {
		// Return an error if the event cannot be retrieved
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create a new event
	var updatedEvent models.Event

	// Re-set the ID and UserID of the event
	updatedEvent.ID = event.ID
	updatedEvent.UserID = event.UserID

	// Bind the request body to the event
	if err := c.ShouldBindJSON(&updatedEvent); err != nil {
		// Return an error if the event cannot be bound
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the event
	err = updatedEvent.Update(convertedId)

	if err != nil {
		// Return an error if the event cannot be updated
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the updated event
	c.JSON(http.StatusOK, gin.H{
		"message": "Your event was succesfully updated!",
		"event":   updatedEvent,
	})
}

func deleteEvent(c *gin.Context) {
	// Get the ID from the URL
	id := c.Param("id")

	// Convert the ID to an integer
	convertedId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		// Return an error if the ID cannot be converted
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	event, err := models.GetEventById(convertedId)

	if event.ID == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if err != nil {
		// Return an error if the event cannot be retrieved
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Delete the event
	err = event.Delete()

	if err != nil {
		// Return an error if the event cannot be deleted
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success message
	c.JSON(http.StatusOK, gin.H{
		"message": "Your event was succesfully deleted!",
	})
}
