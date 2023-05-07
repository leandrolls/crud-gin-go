package main

import (
	"alura-go-gin/controllers"
	"alura-go-gin/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestVerifyStatusCode(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.HelloAluno)
	req, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	if response.Code != http.StatusOK {
		t.Fatalf("Invalid status error received")
	}
}

func TestTAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.GetAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
	fmt.Println(response.Body)
}
