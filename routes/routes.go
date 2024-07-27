package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	/*
		Events
	*/


	// Make a group of routes that require authentication
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", addEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	// Registration routes for each event
	server.POST("/events/:id/register", eventRegister)
	server.DELETE("/events/:id/register", eventUnregister)




	// Define a GET request route
	server.GET("/events", getEvents)

	// GET request route to get a single event
	server.GET("/events/:id", getEvent)

	/*
		Users
	*/

	// Define a POST request route to add a user
	server.POST("/signup", signup)
	server.POST("/login", login)
}
