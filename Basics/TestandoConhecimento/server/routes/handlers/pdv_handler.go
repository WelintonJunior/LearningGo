package handlers

import (
	"example.com/teste/server/internal/app/application"
	"github.com/gin-gonic/gin"
)

type PdvHandlers struct {
	service application.PdvService
}

func NewPdvHandlers(service application.PdvService) *PdvHandlers {
	return &PdvHandlers{service: service}
}

func (h *PdvHandlers) PesquisarPdvPorId(context *gin.Context) {

}
