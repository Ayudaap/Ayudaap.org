package main

import (
	"fmt"
	"log"
	"net/http"

	"Ayudaap.org/cmd/api/handlers"
	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	log.Printf("Nueva peticion IP = %v", r.RemoteAddr)
	fmt.Fprintf(w, "Welcome home!")
}

func main() {

	puerto := 8080
	fmt.Printf("Ejecutando en :%d\n", puerto)

	router := mux.NewRouter().StrictSlash(true)
	api := router.PathPrefix("/api/").Subrouter()
	apiV1 := api.PathPrefix("/v1").Subrouter()
	org := apiV1.PathPrefix("/org").Subrouter()
	org.HandleFunc("/", handlers.ListOrganizaciones).Methods("GET")

	router.HandleFunc("/", homeLink)
	http.ListenAndServe(fmt.Sprintf(":%d", puerto), router)
}
