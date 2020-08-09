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

//DirectorioRepo Modelo de Directorio
var DirectorioRepo *repository.DirectorioRepository

func init() {
	DirectorioRepo = &repository.DirectorioRepository{*repository.GetInstance()}
}

//GetALlDirectorioReq Lista Directorio
func GetALlDirectorioReq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resultados := DirectorioRepo.GetAllDirectorio()

	if len(resultados) <= 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.RespuestaGenerica{Mensaje: "No se encontraron datos a mostrar"})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resultados)
	}
}

//CreateDirectorio Crea una nueva organizacion
func CreateDirectorio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	var directorio models.Directorio

	if err := json.NewDecoder(r.Body).Decode(&directorio); err != nil {
		GetError(err, w)
		return
	}

	directorio.ID = primitive.NewObjectID()

	directorioInsertada := make(chan string)
	defer close(directorioInsertada)

	idInsertado := DirectorioRepo.InsertarDirectorio(directorio)

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

//GetDirectorioById Obtiene una Directorio por ID
func GetDirectorioById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	resultados := DirectorioRepo.GetDirectorioByID(id)

	if resultados == nil {
		json.NewEncoder(w).Encode(models.RespuestaGenerica{Mensaje: "No se encontraron datos a mostrar"})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resultados)
	}
}

// Elimina un directorio
func DeleteDirectorio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	resultados, err := DirectorioRepo.DeleteDirectorio(id)
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
func UpsertDirectorio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	var directorio models.Directorio

	if err := json.NewDecoder(r.Body).Decode(&directorio); err != nil {
		GetError(err, w)
	}

	resultados, err := DirectorioRepo.UpdateDirectorio(&directorio)
	if err != nil {
		GetError(err, w)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Procesado int64 `json:"procesado,omitempty"`
		}{Procesado: resultados})
	}
}
