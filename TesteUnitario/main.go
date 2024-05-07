package main

import (
	"example.com/teste/router"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	router.PostRoute(server)
	router.GetRoute(server)

	server.Run(":3000")
}
