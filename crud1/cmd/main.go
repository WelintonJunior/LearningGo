package main

import (
	"example.com/crud1/internal/adapters/repository"
	db "example.com/crud1/internal/adapters/repository/database"
	"example.com/crud1/internal/app/application"
	"example.com/crud1/routes"
	"example.com/crud1/routes/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	err := db.DB.Ping()

	if err != nil {
		panic(err)
	}

	bookRepo := repository.NewLocalBookRepository()
	bookService := application.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(*bookService)
	routes.BookRoutes(server, *bookHandler)

	server.Run(":3000")
}
