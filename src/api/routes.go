package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetHandler Obtiene el handler principal de la aplicacion
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

	apiOrg := api.PathPrefix("/organizaciones").Subrouter()
	apiOrg.HandleFunc("/", CreateOrganizacion).Methods("POST").Name("createOrganizacion")
	apiOrg.HandleFunc("/", GetALlOrganizaciones).Methods("GET").Name("getAllOrganizacions")
	apiOrg.HandleFunc("/q", GetOrganizacionByQuery).Methods("GET").Name("getOrganizacionByQuery")
	apiOrg.HandleFunc("/{organizacionId}", UpdateOrganizacion).Methods("PUT").Name("updateOrganizacion")
	apiOrg.HandleFunc("/{organizacionId}", GetOrganizacionByID).Methods("GET").Name("getOrganizacion")
	apiOrg.HandleFunc("/{organizacionId}", DeleteOrganizacion).Methods("DELETE").Name("deleteOrganizacion")

	// Direccion
	apiOrg.HandleFunc("/{organizacionId}/direccion", UpdateDireccion).Methods("PUT").Name("updateDireccion")
	apiOrg.HandleFunc("/{organizacionId}/direccion", GetDireccionByOrganizacionID).Methods("GET").Name("getOrganizacionDireccion")

	apiProy := api.PathPrefix("/proyectos").Subrouter()
	apiProy.HandleFunc("/", Createproyecto).Methods("POST").Name("crearProyecto")
	apiProy.HandleFunc("/", GetALlProyectos).Methods("GET").Name("obtenerProyectos")
	apiProy.HandleFunc("/", UpsertProyecto).Methods("PUT").Name("modificarProyecto")
	apiProy.HandleFunc("/{proyectoId}", GetProyectoByID).Methods("GET").Name("getProyectoById")
	apiProy.HandleFunc("/{proyectoId}", DeleteProyecto).Methods("DELETE").Name("borrarProyecto")

	return r
}
