package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibirMensagem(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello Get"})
}
