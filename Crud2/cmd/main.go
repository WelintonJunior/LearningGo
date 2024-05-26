package main

import (
	"example.com/crudbasics/internal/adapters/repository"
	db "example.com/crudbasics/internal/adapters/repository/database"
	"example.com/crudbasics/internal/app/application"
	router "example.com/crudbasics/internal/routes"
	"example.com/crudbasics/internal/routes/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	bookRepo := repository.NewLocalBookRepository()
	bookService := application.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(*bookService)
	router.BookRoutes(server, bookHandler)

	server.Run(":3000")
}
