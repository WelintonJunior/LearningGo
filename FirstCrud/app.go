package main

import (
	"example.com/crud/db"
	"example.com/crud/routes"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	db.InitDB()
	server := gin.Default()

	server.Use(CORSMiddleware())
	routes.ClienteRoutes(server)
	routes.SignupRoutes(server)

	server.Run(":3000")
}
