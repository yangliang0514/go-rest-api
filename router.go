package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yangliang0514/go-rest-api/controllers"
	"github.com/yangliang0514/go-rest-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) *gin.Engine {
	server.GET("/events", middlewares.AuthMiddleware(), controllers.GetAllEvents)
	server.GET("/events/:id", middlewares.AuthMiddleware(), controllers.GetEvent)

	server.PUT("/events/:id", middlewares.AuthMiddleware(), controllers.UpdateEvent)

	server.POST("/events", middlewares.AuthMiddleware(), controllers.CreateEvent)
	server.POST("/events/:id/register", middlewares.AuthMiddleware(), controllers.RegisterUserToEvent)
	server.POST("/events/registered", middlewares.AuthMiddleware(), controllers.ListRegisteredEvents)

	server.DELETE("/events/:id", middlewares.AuthMiddleware(), controllers.DeleteEvent)
	server.DELETE("/events/:id/register", middlewares.AuthMiddleware(), controllers.UnregisterUserFromEvent)

	server.POST("/signup", controllers.Signup)
	server.POST("/login", controllers.Login)

	return server
}
