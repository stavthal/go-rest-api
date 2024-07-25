package routes

import "github.com/gin-gonic/gin"


func RegisterRoutes(server *gin.Engine) {
	// Define a GET request route
	server.GET("/events", getEvents)

	// GET request route to get a single event
	server.GET("/events/:id", getEvent)

	// Define a POST request route to add an event
	server.POST("/events", addEvent)
}