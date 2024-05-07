package router

import (
	"example.com/teste/service"
	"github.com/gin-gonic/gin"
)

func PostRoute(server *gin.Engine) {
	server.POST("/ExibirMensagemComNome", service.ExibirMensagemComNome)
}

func GetRoute(server *gin.Engine) {
	server.GET("/ExibirMensagem", service.ExibirMensagem)
}
