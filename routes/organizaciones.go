package routes

import (
	"log"
	"net/http"
)

type OrganizacionController struct{}

// Lista todas las organizaciones
func (o *OrganizacionController) echoOrg(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola Mundo desde Organizaciones"))
}

func (o *OrganizacionController) Get(w http.ResponseWriter, r *http.Request) {
	log.Printf("Peticion desde %s\n", r.RequestURI)
	w.Write([]byte("Hola desde GET"))
}
