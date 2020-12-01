package common

import (
	"net/http"

	"github.com/gorilla/mux"
)

//GetRouter Obtiene la lista de rutas de la aplicacion
func GetRouter() http.Handler {

	r := mux.NewRouter()
	r.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hola Mundo"))
	}).Methods("GET")

	r.PathPrefix("/api/v1").Subrouter()

	return r
}
