package database

import (
	"log"

	"github.com/danielmaques/go_gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func ConectaBancoDeDados(){
	stringDeConexao := "host=localhost user=root password=root dbname=root"
	DB, err = gorm.Open(postgres.Open(stringDeConexao)) 
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}