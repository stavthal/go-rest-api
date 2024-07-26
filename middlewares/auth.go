package middlewares

import (
	"net/http"

	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	// Get the token from the headers
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		// Return an error if the token is missing
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Verify the token
	userId, err := utils.VerifyToken(token)

	if err != nil {
		// Return an error if the token is invalid
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalid or has expired."})
		return
	}

	// Set the userId in the context
	c.Set("userId", userId)

	// Move to the next request
	c.Next()
}
