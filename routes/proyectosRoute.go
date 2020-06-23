package routes

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"Ayudaap.org/models"
	"Ayudaap.org/repository"
)

// Modelo de Proyectos
var proyRepo *repository.ProyectosRepository

func init() {
	proyRepo = new(repository.ProyectosRepository)
}

// Lista todas las Proyectos
func GetALlProyectosReq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resultados := proyRepo.GetAllProyectos()

	if len(resultados) <= 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.RespuestaGenerica{Mensaje: "No se encontraron datos a mostrar"})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resultados)
	}
}

// Crea una nueva organizacion
func Createproyecto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	var proyecto models.Proyecto

	if err := json.NewDecoder(r.Body).Decode(&proyecto); err != nil {
		GetError(err, w)
	}

	proyecto.ID = primitive.NewObjectID()
	proyecto.Area.ID = primitive.NewObjectID()

	proyectoInsertada := make(chan string)
	defer close(proyectoInsertada)

	idInsertado := proyRepo.InsertProyecto(proyecto)

	if len(idInsertado) <= 0 {
		err := errors.New("No se pudo insertar el objeto")
		GetError(err, w)
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(struct {
			Id string `json:"id,omitempty"`
		}{Id: idInsertado})
	}
}

// Obtiene una proyecto por ID
func GetProyectoById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	resultados := proyRepo.GetProyectoById(id)

	if resultados == nil {
		json.NewEncoder(w).Encode(models.RespuestaGenerica{Mensaje: "No se encontraron datos a mostrar"})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resultados)
	}
}

// Elimina un proyecto
func DeleteProyecto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	resultados, err := proyRepo.DeleteProyecto(id)
	if err != nil {
		GetError(err, w)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Procesado int `json:"procesado,omitempty"`
		}{Procesado: resultados})
	}
}

//Actualiza un objeto
func UpsertProyecto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	var proyecto models.Proyecto

	if err := json.NewDecoder(r.Body).Decode(&proyecto); err != nil {
		GetError(err, w)
	}

	resultados, err := proyRepo.UpdateProyecto(&proyecto)
	if err != nil {
		GetError(err, w)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Procesado int64 `json:"procesado,omitempty"`
		}{Procesado: resultados})
	}
}
