package routes

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"Ayudaap.org/models"
	repo "Ayudaap.org/repository"
)

//GetALlProyectos Lista todas las Proyectos
func GetALlProyectos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resultados := repo.GetAllProyectos()

	if len(resultados) <= 0 {
		GetGenericMessage("No se encontraron datos a mostrar", w)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resultados)
	}
}

//Createproyecto Crea una nueva organizacion
func Createproyecto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	var proyecto models.Proyecto

	if err := json.NewDecoder(r.Body).Decode(&proyecto); err != nil {
		GetError(err, w)
		return
	}

	proyecto.ID = primitive.NewObjectID()
	proyecto.Area.ID = primitive.NewObjectID()

	proyectoInsertada := make(chan string)
	defer close(proyectoInsertada)

	idInsertado := repo.InsertProyecto(proyecto)

	if len(idInsertado) <= 0 {
		err := errors.New("No se pudo insertar el objeto")
		GetError(err, w)
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(struct {
			ID string `json:"id,omitempty"`
		}{ID: idInsertado})
	}
}

//GetProyectoByID Obtiene una proyecto por ID
func GetProyectoByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	resultados := repo.GetProyectoByID(id)

	if resultados == nil {
		GetGenericMessage("No se encontraron datos a mostrar", w)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resultados)
	}
}

//DeleteProyecto Elimina un proyecto
func DeleteProyecto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	resultados, err := repo.DeleteProyecto(id)
	if err != nil {
		GetError(err, w)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Procesado int `json:"procesado,omitempty"`
		}{Procesado: resultados})
	}
}

//UpsertProyecto Actualiza un objeto
func UpsertProyecto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	var proyecto models.Proyecto

	if err := json.NewDecoder(r.Body).Decode(&proyecto); err != nil {
		GetError(err, w)
	}

	resultados, err := repo.UpdateProyecto(&proyecto)
	if err != nil {
		GetError(err, w)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Procesado int64 `json:"procesado,omitempty"`
		}{Procesado: resultados})
	}
}
