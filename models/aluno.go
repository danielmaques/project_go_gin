package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	Nome string `json:"nome"`
	CPF string `json:"cpf"`
}

var Alunos [] Aluno