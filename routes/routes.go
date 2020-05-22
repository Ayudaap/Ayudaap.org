package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"Ayudaap.org/models"
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

	apiOrg := api.PathPrefix("/organizacion").Subrouter()
	apiOrg.HandleFunc("/", GetALlOrganizacionesReq).Methods("GET").Name("getAllOrganizaciones")

	//TODO: Borrarla en productivo
	apiOrg.HandleFunc("/inicializarDemo", InicializarOrganizaciones).Methods("GET").Name("inicializarOrganizaciones")

	return r
}

// GetError : This is helper function to prepare error model.
// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
func GetError(err error, w http.ResponseWriter) {

	log.Fatal(err.Error())
	var response = models.ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
