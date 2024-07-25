package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func signup (c *gin.Context) {
	// Create a new user
	var user models.User
	
	// Bind the request body to the user
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Save the user
	if err := user.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(201, gin.H{"status": "User successfully created"})
}

func login(c *gin.Context) {
	// Create a new user
	var user models.User
	
	// Bind the request body to the user
	err := c.ShouldBindJSON(&user) 
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Authenticate the user
	err = user.Authenticate()
	
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"status": "User successfully authenticated"})
}