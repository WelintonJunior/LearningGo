package router

import (
	"example.com/crudbasics/internal/routes/handlers"
	"github.com/gin-gonic/gin"
)

func BookRoutes(server *gin.Engine, handlers *handlers.BookHandler) {
	server.POST("/Book/New", handlers.NewBook)
}
