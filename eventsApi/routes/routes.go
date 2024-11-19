package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	//Events routes

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", CreateEventHandler)
	authenticated.PUT("/events/:id", UpdateEvent)
	authenticated.DELETE("/events/:id", DeleteEvent)

	server.GET("/events", GetEventsHandler)
	server.GET("/events/:id", GetSingleEventHandler)
	//User routes
	server.POST("/signup", SignUp)
	server.POST("/login", Login)
	server.GET("/users", ListUsers)
	//Registration routes
	authenticated.POST("/events/:id/register", RegisterForEvent)
	authenticated.DELETE("/events/:id/register", CancelRegistration)

}
