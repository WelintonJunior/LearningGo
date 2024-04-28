package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	server.GET("/events", getEvents)

	server.Run(":3000")
	// http.ListenAndServe(":3000", nil)
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"name": "Welinton",
	})
}
