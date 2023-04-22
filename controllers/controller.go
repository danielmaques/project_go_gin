package controllers

import (
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