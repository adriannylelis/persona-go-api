package main

import (
	"fmt"

	database "github.com/adriannylelis/persona-go-api/Database"
	"github.com/adriannylelis/persona-go-api/models"
	"github.com/adriannylelis/persona-go-api/routes"
)


func main()  {
models.Personalidades =[]models.Personalidade{
	{Nome: "Nome 1", Descricao: "Descricao 1"},
}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando aplicacao...")
	routes.HandleRequest()
}