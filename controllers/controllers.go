package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/adriannylelis/persona-go-api/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Home Page")
	
}

func GetPersonalidades(w http.ResponseWriter, r *http.Request)  {
	json.NewEncoder(w).Encode(models.Personalidades)
	
}

func GetPersonalidade(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id:= vars["id"]

	for _, personalidade := range models.Personalidades{
		if strconv.Itoa(personalidade.ID) == id {
			json.NewEncoder(w).Encode(personalidade)
			return
		}
	}
}