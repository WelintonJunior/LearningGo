package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Mensagem struct {
	Name string
}

func ExibirMensagemComNome(context *gin.Context) {
	var nome Mensagem
	if err := context.ShouldBindJSON(&nome); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Erro ao receber nome"})
		return
	}

	content := fmt.Sprintf("Hello %v", nome.Name)

	context.JSON(http.StatusOK, gin.H{"message": content})
}
