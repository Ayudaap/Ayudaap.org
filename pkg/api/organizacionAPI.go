package api

import (
	"encoding/json"
	"log"
	"net/http"

	"Ayudaap.org/src/entities"
	"Ayudaap.org/src/models"
	"github.com/gorilla/mux"
)

//OrganizacionAPI API de organizacion
type OrganizacionAPI struct{}

//GetALlorganizaciones Lista todas las organizaciones
func (o OrganizacionAPI) GetALlorganizaciones(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resultados, err := models.OrganizacionModel{}.FindAll()

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	if len(resultados) <= 0 {
		respondWithError(w, http.StatusNotFound, "No se encontraron datos a procesar")

	} else {
		w.WriteHeader(http.StatusOK)
		respondWithJSON(w, http.StatusOK, resultados)
	}
}

func (o OrganizacionAPI) GetOrganizacionById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	resultados, err := models.OrganizacionModel{}.FindByID(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, resultados)
	}
}

//CreateOrganizacion Crea un nuevo Organizacion
func (o OrganizacionAPI) CreateOrganizacion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	var organizacion entities.Organizacion

	if err := json.NewDecoder(r.Body).Decode(&organizacion); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	resultados, err := models.OrganizacionModel{}.InsertOne(organizacion)
	if err != nil {
		log.Fatal(err.Error())
		respondWithError(w, http.StatusBadRequest, "No se pudo procesar la peticion")
	} else {
		respondWithJSON(w, http.StatusCreated, resultados)
	}
}

// func (o OrganizacionAPI) GetOrganizacionById(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	defer r.Body.Close()
// }
