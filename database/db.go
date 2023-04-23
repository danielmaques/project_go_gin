package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
    DB *gorm.DB
    err error
)

func Connect(connectionString string) (*gorm.DB, error) {
    DB, err = gorm.Open(postgres.Open(connectionString))
    if err != nil {
        log.Panic("Erro ao conectar com o banco de dados")
        return nil, err
    }
    return DB, nil
}
