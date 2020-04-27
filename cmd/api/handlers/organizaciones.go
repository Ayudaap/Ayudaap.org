package handlers

import (
	"net/http"
)

// Lista todas las organizaciones
func ListOrganizaciones(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola Mundo"))
}
