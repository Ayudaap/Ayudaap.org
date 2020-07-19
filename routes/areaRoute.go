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

// Modelo de Areas
var areaRepo *repository.GenericRepository

func init() {
	areaRepo = new(repository.GenericRepository)
}

// Lista todas las areas
func GetALlAreasReq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resultados := areaRepo.GetAll()

	if len(resultados) <= 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.RespuestaGenerica{Mensaje: "No se encontraron datos a mostrar"})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resultados)
	}
}

// Crea una nueva organizacion
func CreateOrganizacion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	var organizacion models.Organizacion

	if err := json.NewDecoder(r.Body).Decode(&organizacion); err != nil {
		GetError(err, w)
	}

	organizacion.ID = primitive.NewObjectID()
	organizacion.Domicilio.ID = primitive.NewObjectID()
	for _, dir := range organizacion.Domicilio.Directorio {
		dir.ID = primitive.NewObjectID()
	}

	organizacionInsertada := make(chan string)
	defer close(organizacionInsertada)

	idInsertado := areaRepo.InsertOrganizacion(organizacion)

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

// Obtiene una organizacion por ID
func GetOrganizacionById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	resultados := areaRepo.GetOrganizacionById(id)

	if resultados == nil {
		json.NewEncoder(w).Encode(models.RespuestaGenerica{Mensaje: "No se encontraron datos a mostrar"})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resultados)
	}
}

// Elimina una organizacion
func DeleteOrganizacion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	resultados, err := areaRepo.DeleteOrganizacion(id)
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
func UpsertOrganizacion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	var organizacion models.Organizacion

	if err := json.NewDecoder(r.Body).Decode(&organizacion); err != nil {
		GetError(err, w)
	}

	resultados, err := areaRepo.UpdateOrganizacion(&organizacion)
	if err != nil {
		GetError(err, w)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Procesado int64 `json:"procesado,omitempty"`
		}{Procesado: resultados})
	}
}
