package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yangliang0514/go-rest-api/database"
)

func main() {
	server := gin.Default()
	RegisterRoutes(server)
	database.InitDB()
	server.Run(":8080")
}
