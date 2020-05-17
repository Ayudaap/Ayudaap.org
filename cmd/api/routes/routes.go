package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Obtiene el handler principal de la aplicacion
func GetHandler() http.Handler {
	r := mux.NewRouter()
	r.StrictSlash(true)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Peticion en %v desde %v", r.RequestURI, r.RemoteAddr)
		w.Write([]byte("Hola desde default"))
	})

	// TODO: Cambiar la version de API desde configuracion
	api := r.PathPrefix("/api/v1").Subrouter()
	api.StrictSlash(true)

	organizacionRouter := new(OrganizacionController)
	apiOrg := api.PathPrefix("/organizacion").Subrouter()
	apiOrg.HandleFunc("/", organizacionRouter.Get).Methods("GET")
	apiOrg.HandleFunc("/echo", organizacionRouter.echoOrg)

	return r
}
