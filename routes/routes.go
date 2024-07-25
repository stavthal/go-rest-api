package routes

import "github.com/gin-gonic/gin"


func RegisterRoutes(server *gin.Engine) {

	/*
		Events
	*/

	// Define a GET request route
	server.GET("/events", getEvents)

	// GET request route to get a single event
	server.GET("/events/:id", getEvent)

	// Define a POST request route to add an event
	server.POST("/events", addEvent)

	// PUT request route to update an event
	server.PUT("/events/:id", updateEvent)

	// DELETE request route to delete an event
	server.DELETE("/events/:id", deleteEvent)


	/*
		Users
	*/

	// Define a POST request route to add a user
	server.POST("/signup", signup)
	server.POST("/login", login)
}