package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yangliang0514/go-rest-api/controllers"
)

func RegisterRoutes(server *gin.Engine) *gin.Engine {
	server.GET("/events", controllers.GetAllEvents)
	server.GET("/events/:id", controllers.GetEvent)
	server.POST("/events", controllers.CreateEvent)
	server.PUT("/events/:id", controllers.UpdateEvent)
	server.DELETE("/events/:id", controllers.DeleteEvent)

	return server
}
