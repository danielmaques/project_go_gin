package models

type Aluno struct {
	Nome string `json:"nome"`
	CPF string `json:"cpf"`
}

var Alunos [] Aluno