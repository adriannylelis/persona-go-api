package routes

import (
	"log"
	"net/http"

	"github.com/adriannylelis/persona-go-api/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest()  {

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalidades", controllers.GetPersonalidades).Methods("GET")
	r.HandleFunc("/personalidades/{id}", controllers.GetPersonalidade).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}