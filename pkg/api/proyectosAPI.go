package api

import (
	"encoding/json"
	"log"
	"net/http"

	"Ayudaap.org/src/entities"
	"Ayudaap.org/src/models"
)

//ProyectoAPI Api de proyectos
type ProyectoAPI struct{}

//GetALlProyectos Lista todas las Proyectos
func (p ProyectoAPI) GetALlProyectos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resultados, err := models.ProyectoModel{}.FindAll()

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

//CreateProyecto Crea un nuevo proyecto
func (p ProyectoAPI) CreateProyecto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	var proyecto entities.Proyecto

	if err := json.NewDecoder(r.Body).Decode(&proyecto); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	resultados, err := models.ProyectoModel{}.InsertOne(proyecto)
	if err != nil {
		log.Fatal(err.Error())
		respondWithError(w, http.StatusBadRequest, "No se pudo procesar la peticion")
	} else {
		respondWithJSON(w, http.StatusCreated, resultados)
	}
}
