package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"Ayudaap.org/models"
	"github.com/gorilla/mux"
)

//GetHandler Obtiene el handler principal de la aplicacion
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
	apiOrg.HandleFunc("/", CreateOrganizacion).Methods("POST").Name("crearOrganizacion")
	apiOrg.HandleFunc("/", GetALlOrganizaciones).Methods("GET").Name("obtenerOrganizaciones")
	apiOrg.HandleFunc("/", UpsertOrganizacion).Methods("PUT").Name("modificarOrganizacion")
	apiOrg.HandleFunc("/{id}", GetOrganizacionByID).Methods("GET").Name("getOrganizacionById")
	apiOrg.HandleFunc("/{id}", DeleteOrganizacion).Methods("DELETE").Name("borrarOrganizacion")
	apiOrg.HandleFunc("/{id}/direccion", GetDireccionByOrganizacionID).Methods("GET").Name("getOrganizacionByOrganizacionID")

	apiProy := api.PathPrefix("/proyecto").Subrouter()
	apiProy.HandleFunc("/", Createproyecto).Methods("POST").Name("crearProyecto")
	apiProy.HandleFunc("/", GetALlProyectos).Methods("GET").Name("obtenerProyectos")
	apiProy.HandleFunc("/", UpsertProyecto).Methods("PUT").Name("modificarProyecto")
	apiProy.HandleFunc("/{id}", GetProyectoByID).Methods("GET").Name("getProyectoById")
	apiProy.HandleFunc("/{id}", DeleteProyecto).Methods("DELETE").Name("borrarProyecto")

	return r
}

// GetError : This is helper function to prepare error model.
// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
func GetError(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	var response = models.ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}

// GetGenericError Genera un mensaje generico de error
func GetGenericMessage(mensaje string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.RespuestaGenerica{Mensaje: mensaje})
}
