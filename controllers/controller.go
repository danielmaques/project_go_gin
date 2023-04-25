package controllers

import (
	"net/http"

	"github.com/danielmaques/go_gin/database"
	"github.com/danielmaques/go_gin/models"
	"github.com/gin-gonic/gin"
)

func ExibirAlunos(c *gin.Context){
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
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

func BuscarID(c *gin.Context){
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	c.JSON(http.StatusOK, aluno)
}
