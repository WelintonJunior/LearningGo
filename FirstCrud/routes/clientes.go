package routes

import (
	"net/http"
	"strconv"

	"example.com/crud/models"
	"github.com/gin-gonic/gin"
)

func CreateCliente(context *gin.Context) {
	var cliente models.Cliente

	if err := context.ShouldBindJSON(&cliente); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Falha ao transformar os dados em JSON", "error": err})
		return
	}

	if err := models.Cliente.New(cliente); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Falha ao criar cliente", "error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Sucesso!"})
}

func ReadClientes(context *gin.Context) {
	clientes, err := models.List()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Erro ao coletar dados", "error": err})
		return
	}

	context.JSON(http.StatusOK, clientes)
}

func ReadCliente(context *gin.Context) {
	clienteId, _ := strconv.ParseInt(context.Param("id"), 10, 64)
	cliente, err := models.Read(clienteId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Erro ao coletar dados", "error": err})
		return
	}

	context.JSON(http.StatusOK, cliente)
}

func UpdateCliente(context *gin.Context) {
	var cliente models.Cliente

	if err := context.ShouldBindJSON(&cliente); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Erro ao transformar dados", "error": err})
		return
	}

	if err := cliente.Update(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao atualizar dados", "error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Sucesso!"})
}

func DeleteCliente(context *gin.Context) {
	var cliente models.Cliente

	if err := context.ShouldBindJSON(&cliente); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Erro ao ler id", "error": err})
		return
	}

	if err := cliente.Delete(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Erro Deletar Cliente", "error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Sucesso!"})
}
