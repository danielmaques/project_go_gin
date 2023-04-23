package controllers

import (
	"net/http"

	"github.com/danielmaques/go_gin/database"
	"github.com/danielmaques/go_gin/models"
	"github.com/gin-gonic/gin"
)

func ExibirAlunos(c *gin.Context){
	c.JSON(200, models.Alunos)
}

func BemVindo(c *gin.Context){
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API": "Bem vindo" + nome,
	})
}

func CriarNovoAluno(c *gin.Context){
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
			return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}
