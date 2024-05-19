package main

import (
	db "example.com/teste/server/database"
	"example.com/teste/server/internal/adapteres/repository"
	"example.com/teste/server/internal/app/application"
	"example.com/teste/server/routes"
	"example.com/teste/server/routes/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	repoPdv := repository.NewLocalPdvRepository()
	servicePdv := application.NewPdvService(repoPdv)
	handlerPdv := handlers.NewPdvHandlers(*servicePdv)
	routes.PdvRoutes(server, handlerPdv)

	server.Run(":3000")
}
