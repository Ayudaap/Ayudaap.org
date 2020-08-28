package organizaciones

import (
	"encoding/json"
	"net/http"

	rp "Ayudaap.org/pkg/repository"
)

// Nombre de la tabla de organizaciones
const organizacionCollection string = "organizaciones"

// Obtiene un acceso a la base de datos
func getConexion() *rp.MongoRepository {
	db := rp.GetInstance()

	return db
}

//GetALlOrganizacionesReq Lista todas las organizaciones
func GetALlOrganizacionesReq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resultados := GetAllOrganizaciones()

	if len(resultados) <= 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(struct {
			Mensaje string `json:"mensaje,omitempty"`
		}{Mensaje: "No se encontraron datos a mostrar"})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resultados)
	}
}
