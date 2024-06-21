package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adriannylelis/persona-go-api/database"
	"github.com/adriannylelis/persona-go-api/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Home Page")
	
}

func GetPersonalidades(w http.ResponseWriter, r *http.Request)  {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
	
}

func GetPersonalidade(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id:= vars["id"]

	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}