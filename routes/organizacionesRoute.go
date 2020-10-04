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

//GetALlOrganizaciones Lista todas las organizaciones
func GetALlOrganizaciones(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resultados := repo.GetAllOrganizaciones()

	if len(resultados) <= 0 {
		GetGenericMessage("No se encontraron datos a mostrar", w)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resultados)
	}
}

//GetDireccionByOrganizacionID Obtiene la direccion por el Id de una organizacion
func GetDireccionByOrganizacionID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	resultados, error := repo.GetDireccionByOrganizacionID(id)

	if error != nil {
		GetGenericMessage("No se encontraron datos a mostrar", w)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resultados)
	}
}

//CreateOrganizacion Crea una nueva organizacion
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

	idInsertado := repo.InsertOrganizacion(organizacion)

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

//GetOrganizacionByID Obtiene una organizacion por ID
func GetOrganizacionByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	resultados := repo.GetOrganizacionByID(id)

	if resultados == nil {
		json.NewEncoder(w).Encode(models.RespuestaGenerica{Mensaje: "No se encontraron datos a mostrar"})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resultados)
	}
}

//DeleteOrganizacion Elimina una organizacion
func DeleteOrganizacion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	resultados, err := repo.DeleteOrganizacion(id)
	if err != nil {
		GetError(err, w)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Procesado int `json:"procesado,omitempty"`
		}{Procesado: resultados})
	}
}

//UpsertOrganizacion Actualiza un objeto
func UpsertOrganizacion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	var organizacion models.Organizacion

	if err := json.NewDecoder(r.Body).Decode(&organizacion); err != nil {
		GetError(err, w)
	}

	resultados, err := repo.UpdateOrganizacion(&organizacion)
	if err != nil {
		GetError(err, w)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Procesado int64 `json:"procesado,omitempty"`
		}{Procesado: resultados})
	}
}
