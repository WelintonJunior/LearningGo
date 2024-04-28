package main

import (
	"net/http"

	"example.com/api/models"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":3000")
	// http.ListenAndServe(":3000", nil)
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err})
		return
	}

	event.ID = 1
	event.UserID = 1

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
