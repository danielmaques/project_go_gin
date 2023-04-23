package main

import (
	"github.com/danielmaques/go_gin/database"
	"github.com/danielmaques/go_gin/models"
	"github.com/danielmaques/go_gin/routes"
)

func main(){
	database.ConectaBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Julho", CPF: "390.407.090-84"},
		{Nome: "Ana", CPF: "014.051.990-40"},
	}
	routes.HandleRequests()
}