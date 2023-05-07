package routes

import (
	"alura-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/alunos", controllers.GetAlunos)
	r.GET("/alunos/:id", controllers.GetAluno)
	r.GET("/alunos/cpf/:cpf", controllers.GetAlunoByCPF)
	r.DELETE("/alunos/:id", controllers.DeleteAluno)
	r.PATCH("/alunos/:id", controllers.EditAluno)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/:nome", controllers.HelloAluno)
	r.GET("/index", controllers.ExibePaginaIndex)
	r.NoRoute(controllers.RotaNaoEncontrada)
	r.Run(":5000")
}
