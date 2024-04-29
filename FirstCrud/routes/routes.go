package routes

import "github.com/gin-gonic/gin"

func ClienteRoutes(server *gin.Engine) {
	server.POST("/CreateCliente", CreateCliente)
	server.GET("/ReadClientes", ReadClientes)
	server.GET("/ReadCliente/:id", ReadCliente)
	server.PUT("/UpdateCliente", UpdateCliente)
	server.DELETE("/DeleteCliente", DeleteCliente)
}

func SignupRoutes(server *gin.Engine) {
	server.POST("/Signup", Signup)
	server.POST("/Login", Login)
}
