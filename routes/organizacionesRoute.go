package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"Ayudaap.org/models"
)

var organizaciones []models.Organizacion

func init() {
	organizaciones = append(organizaciones, models.Organizacion{
		Tipo: models.OrganizacionNoGubernamental,
		Domicilio: models.Direccion{
			Calle:          "Siempre Viva",
			NumeroExterior: "123",
			CodigoPostal:   "66001",
			Colonia:        "Springfield",
			Estado:         "Nuevo Leon",
		},
		Nombre:             "Simpsonia",
		RepresentanteLegal: "Homero J. Simpson",
	})
}

// Lista todas las organizaciones

func GetALlOrganizaciones(w http.ResponseWriter, r *http.Request) {
	log.Printf("Peticion desde %s\n", r.RequestURI)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(organizaciones)
}
