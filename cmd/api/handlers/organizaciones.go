package handlers

import (
	"log"
	"net/http"
)

type Organizacion struct{}

// Lista todas las organizaciones
func (o *Organizacion) echoOrg(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola Mundo desde Organizaciones"))
}

func (o *Organizacion) Get(w http.ResponseWriter, r *http.Request) {
	log.Printf("Peticion desde %s\n", r.RequestURI)
	w.Write([]byte("Hola desde GET"))
}
