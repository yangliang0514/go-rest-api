package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yangliang0514/go-rest-api/database"
	"github.com/yangliang0514/go-rest-api/router"
)

func main() {
	gin.DisableConsoleColor() // Disable Gin's default colored output
	server := router.RegisterRoutes(gin.Default())
	database.InitDB()
	server.Run(":8080")
}
