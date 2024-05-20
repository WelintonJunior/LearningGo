package handlers

import (
	"net/http"

	"example.com/crud1/internal/app/application"
	"example.com/crud1/internal/app/domain"
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
		c.JSON(http.StatusBadRequest, gin.H{"message": "Erro ao receber dados", "error": true})
		return
	}

	if err := h.service.NewBook(b); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao cadastrar livro", "error": true})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sucesso!", "error": false})
}

func (h *BookHandler) ListBooks(c *gin.Context) {
	books, err := h.service.ListBooks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao cadastrar livro", "error": true})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) ListBooksNoRoutine(c *gin.Context) {
	books, err := h.service.ListBooksNoRoutine()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao cadastrar livro", "error": true})
		return
	}

	c.JSON(http.StatusOK, books)
}
