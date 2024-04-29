package routes

import (
	"net/http"

	"example.com/crud/models"
	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context) {
	var cliente models.Cliente

	if err := context.ShouldBindJSON(&cliente); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Falha ao trasnformar dados", "error": err})
		return
	}

	if err := cliente.New(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Falha ao criar cliente", "error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Sucesso!"})

}

func Login(context *gin.Context) {
	var cliente models.Cliente

	if err := context.ShouldBindJSON(&cliente); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Falha ao trasnformar dados", "error": err})
		return
	}

	if err := cliente.ValidateCredentials(); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"messsage": "Erro ao fazer login", "error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"messsage": "Sucesso!"})

}
