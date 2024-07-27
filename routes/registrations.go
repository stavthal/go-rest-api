package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func eventRegister(c *gin.Context) {
	userID := c.GetInt64("userId")
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)

	// Check if the event exists
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Get the event from its ID
	event, err := models.GetEventById(eventID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event"})
		return
	}

	// Register the user for the event
	err = event.RegisterUser(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user for event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered for event",
		"event":   event,
	})
}


func eventUnregister(c *gin.Context) {
	userID := c.GetInt64("userId")
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)

	// Check if the event exists
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Get the event from its ID
	event, err := models.GetEventById(eventID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event"})
		return
	}

	// Unregister the user for the event
	err = event.UnregisterUser(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not unregister user for event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User unregistered for event",
		"event":   event,
	})
}