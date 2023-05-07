package controllers

import (
	"alura-go-gin/database"
	"alura-go-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}

func GetAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Find(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func GetAlunoByCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)
	if aluno.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func HelloAluno(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(http.StatusOK, gin.H{
		"API diz": "E ai " + nome,
	})
}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.Validate(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func DeleteAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{
		"DEU CERTO": "aluno deletado"})
}

func EditAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Find(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.Validate(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

func ExibePaginaIndex(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

func RotaNaoEncontrada(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
