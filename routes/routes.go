package routes

import (
	"log"
	"net/http"

	"github.com/adriannylelis/persona-go-api/controllers"
)

func HandleRequest()  {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/personalidades", controllers.GetPersonalidades)
	log.Fatal(http.ListenAndServe(":8080", nil))
}