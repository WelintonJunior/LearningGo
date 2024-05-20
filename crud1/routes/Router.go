package routes

import (
	"example.com/crud1/routes/handlers"
	"github.com/gin-gonic/gin"
)

func BookRoutes(server *gin.Engine, handler handlers.BookHandler) {
	server.POST("/Book/New", handler.NewBook)
	server.GET("/Book/List", handler.ListBooks)
	server.GET("/Book/ListNo", handler.ListBooksNoRoutine)
}
