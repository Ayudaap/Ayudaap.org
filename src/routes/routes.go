package routes

import (
	"net/http"

	"Ayudaap.org/pkg/api"
	"github.com/gorilla/mux"
)

//GetRouter Obtiene la lista de rutas de la aplicacion
func GetRouter() http.Handler {

	r := mux.NewRouter()
	r.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hola Mundo"))
	}).Methods("GET")

	apiRoute := r.PathPrefix("/api/v1").Subrouter()
	apiRoute.StrictSlash(true)

	//apiProy Rutas para los proyectos
	apiProy := apiRoute.PathPrefix("/proyectos").Subrouter()
	apiProy.HandleFunc("/", api.ProyectoAPI{}.GetALlProyectos).Name("getALlProyectos").Methods("GET")
	apiProy.HandleFunc("/purgar", api.ProyectoAPI{}.PurgarCollection).Name("purgarProyectos").Methods("GET")
	apiProy.HandleFunc("/{id}", api.ProyectoAPI{}.GetProyectoByID).Name("getProyecto").Methods("GET")
	apiProy.HandleFunc("/", api.ProyectoAPI{}.CreateProyecto).Name("crearProyecto").Methods("POST")

	//apiOrg Rutas para las Organizaciones
	apiOrg := apiRoute.PathPrefix("/organizaciones").Subrouter()
	apiOrg.HandleFunc("/", api.OrganizacionAPI{}.GetALlorganizaciones).Name("getALlOrganizaciones").Methods("GET")
	apiOrg.HandleFunc("/", api.OrganizacionAPI{}.CreateOrganizacion).Name("crearProyecto").Methods("POST")

	return r
}
