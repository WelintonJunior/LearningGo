package main

import (
	"fmt"
	"net/http"

	"example.com/crud/db"
	"example.com/crud/models"
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

	server.POST("/CreateCliente", CreateCliente)
	server.GET("/ReadCliente", ReadCliente)
	server.PUT("/UpdateCliente", UpdateCliente)
	server.DELETE("/DeleteCliente", DeleteCliente)

	server.Run(":3000")
}

func CreateCliente(context *gin.Context) {
	var cliente models.Cliente

	if err := context.ShouldBindJSON(&cliente); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Falha ao transformar os dados em JSON", "error": err})
		return
	}

	fmt.Println(cliente)

	if err := models.Cliente.New(cliente); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Falha ao criar cliente", "error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Sucesso!"})
}

func ReadCliente(context *gin.Context) {
	clientes, err := models.List()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Erro ao coletar dados", "error": err})
		return
	}

	context.JSON(http.StatusOK, clientes)
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
