package rutas

import (
	"log"
	"net/http"

	"Ayudaap.org/pkg/organizaciones"
	"github.com/gorilla/mux"
)

// GetRutas Obtiene la lista de rutas
func GetRutas() http.Handler {

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
	apiOrg.HandleFunc("/", organizaciones.GetALlOrganizacionesReq).Methods("GET").Name("obtenerOrganizaciones")

	return r
}
