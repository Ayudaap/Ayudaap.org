package api

import (
	"encoding/json"
	"log"
	"net/http"

	"Ayudaap.org/src/entities"
	"Ayudaap.org/src/models"
	"github.com/gorilla/mux"
)

//ProyectoAPI Api de proyectos
type ProyectoAPI struct{}

//CreateProyecto Crea un nuevo proyecto
func (p ProyectoAPI) CreateProyecto(w http.ResponseWriter, r *http.Request) {
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

//GetALlProyectos Lista todas las Proyectos
func (p ProyectoAPI) GetALlProyectos(w http.ResponseWriter, r *http.Request) {

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

//GetProyectoByID Obtiene un proyecto por el ID
func (p ProyectoAPI) GetProyectoByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	resultados, err := models.ProyectoModel{}.FindByID(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, resultados)
	}
}

// //PurgarCollection Purga la coleccion de la base de datos
// func (p ProyectoAPI) PurgarCollection(w http.ResponseWriter, r *http.Request) {

// 	err := models.ProyectoModel{}.Purge()

// 	if err != nil {
// 		respondWithJSON(w, http.StatusGone, "")
// 	} else {
// 		respondWithError(w, http.StatusBadRequest, err.Error())
// 	}
// }
