package main

import (
	"github.com/danielmaques/go_gin/database"
	"github.com/danielmaques/go_gin/routes"
)

func main(){
	database.ConectaBancoDeDados()
	routes.HandleRequests()
}