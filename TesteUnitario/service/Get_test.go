package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGet(t *testing.T) {
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)

	ExibirMensagem(context)

	if recorder.Code != http.StatusOK {
		t.Errorf("Esperado status %d; recebeu %d", http.StatusOK, recorder.Code)
	}

	expected := `{"message":"Hello Get"}`
	if body := recorder.Body.String(); body != expected {
		t.Errorf("Esperado corpo '%s'; recebeu '%s'", expected, body)
	}
}
