package routes

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

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
		GetGenericMessage("No se encontraron datos a mostrar", http.StatusOK, w)
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

	idInsertado, err := repo.InsertOrganizacion(organizacion)

	if len(idInsertado) <= 0 || err != nil {
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
	id := mux.Vars(r)["organizacionId"]
	resultados := repo.GetOrganizacionByID(id)

	if resultados == nil {
		json.NewEncoder(w).Encode(models.RespuestaGenerica{Mensaje: "No se encontraron datos a mostrar"})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resultados)
	}
}

//GetOrganizacionByQuery Obtiene los parametros de la consulta
func GetOrganizacionByQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		json.NewEncoder(w).Encode(models.RespuestaGenerica{Mensaje: "No se encontraron datos a mostrar"})
		return
	}

	parametros := make(map[string]string)

	parametrosQuery := strings.Split(keys[0], ",")

	for _, v := range parametrosQuery {
		parametro := strings.Split(v, "=")
		parametros[parametro[0]] = parametro[1]
	}

	resultados, err := repo.GetOrganizacionByQuery(parametros)
	if err != nil {
		err := errors.New("No se pudo localizar el objeto")
		GetError(err, w)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resultados)
	}
}

//DeleteOrganizacion Elimina una organizacion
func DeleteOrganizacion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["organizacionId"]

	resultados, err := repo.DeleteOrganizacion(id)
	if err != nil {
		GetError(err, w)
	} else {
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(struct {
			Procesado int `json:"procesado,omitempty"`
		}{Procesado: resultados})
	}
}

//UpdateOrganizacion Actualiza un objeto
func UpdateOrganizacion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["organizacionId"]

	defer r.Body.Close()

	var organizacion models.Organizacion

	if err := json.NewDecoder(r.Body).Decode(&organizacion); err != nil {
		GetError(err, w)
	}

	resultados, err := repo.UpdateOrganizacion(id, &organizacion)
	if err != nil {
		GetError(err, w)
	} else {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(struct {
			Procesado int64 `json:"procesado,omitempty"`
		}{Procesado: resultados})
	}
}

// Información de direccion

//GetDireccionByOrganizacionID Obtiene la direccion por el Id de una organizacion
func GetDireccionByOrganizacionID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["organizacionId"]

	resultados, error := repo.GetDireccionByOrganizacionID(id)

	if error != nil {
		GetGenericMessage("No se encontraron datos a mostrar", http.StatusOK, w)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resultados)
	}
}

// UpdateDireccion Actualiza la direccion de la organizacion
func UpdateDireccion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["organizacionId"]

	var direccion *models.Direccion

	if err := json.NewDecoder(r.Body).Decode(&direccion); err != nil {
		GetError(err, w)
	}

	resultados, err := repo.UpdateDireccion(id, direccion)
	if err != nil {
		GetGenericMessage("No pudo ejecutar la operacion", http.StatusMethodNotAllowed, w)
	} else {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(struct {
			Procesado int64 `json:"procesado,omitempty"`
		}{Procesado: resultados})
	}

}
