package routes

import (
	"github.com/danielmaques/go_gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests(){
	r := gin.Default()
	r.GET("/alunos", controllers.ExibirAlunos)
	r.GET("/:nome", controllers.BemVindo)
	r.GET("/alunos/:id", controllers.BuscarID)
	r.POST("/alunos", controllers.CriarNovoAluno)
	r.Run()
}