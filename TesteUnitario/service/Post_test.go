package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPost(t *testing.T) {
	p := Mensagem{
		Name: "Welinton",
	}

	jsonData, err := json.Marshal(p)

	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest(http.MethodPost, "/ExibirMensagemComNome", bytes.NewReader(jsonData))

	if err != nil {
		t.Error(err)
	}

	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = req

	ExibirMensagemComNome(context)

	if recorder.Code != http.StatusOK {
		t.Errorf("Esperado Status: %d, Recebeu status: %d", http.StatusOK, recorder.Code)
	}
}
