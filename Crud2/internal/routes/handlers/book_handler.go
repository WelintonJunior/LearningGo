package handlers

import (
	"fmt"
	"net/http"

	"example.com/crudbasics/internal/app/application"
	"example.com/crudbasics/internal/domain"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	service application.BookService
}

func NewBookHandler(service application.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) NewBook(c *gin.Context) {
	var b domain.Book
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Erro ao receber dados"})
		return
	}

	if err := h.service.New(b); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao criar livro"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sucesso"})
}
