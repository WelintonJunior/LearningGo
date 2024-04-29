package main

import (
	"net/http"

	"example.com/api/db"
	"example.com/api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":3000")
}

// FUNÇÃO PARA O ENDPOINT
func getEvents(context *gin.Context) {
	//RECEBE OS EVENTOS DO SELECT
	events, err := models.GetAllEvents()

	//DEVOLVE UMA MENSAGEM DE ERRO
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	//DEVOLVER OS EVENTOS
	context.JSON(http.StatusOK, events)
}

// FUNÇÃO PARA O ENDPOINT
func createEvent(context *gin.Context) {
	var event models.Event
	//TRANSFORMA O STRUCT EVENT EM JSON
	err := context.ShouldBindJSON(&event)

	//DEVOLVE UM ERRO SE NAO CONSEGUIR
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err})
		return
	}

	event.ID = 1
	event.UserID = 1

	//CHAMA A FUNÇÂO SAVE DO PACOTE MODELS PARA DAR INSERT
	err = models.Event.Save(event)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event"})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
