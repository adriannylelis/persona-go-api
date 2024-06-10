package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adriannylelis/persona-go-api/models"
)

func Home(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Home Page")
	
}

func GetPersonalidades(w http.ResponseWriter, r *http.Request)  {
	json.NewEncoder(w).Encode(models.Personalidades)
	
}