package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	RegisterRoutes(server)
	InitDB()
	server.Run(":8080")
}
