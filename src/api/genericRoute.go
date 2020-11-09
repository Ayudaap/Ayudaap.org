package api

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

//RespondWithJSON Responde con un mensaje de error genérico
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
	return
}

//RespondWithError Responde con un mensaje de error
func RespondWithError(w http.ResponseWriter, code int, msj string) {
	RespondWithJSON(w, code, map[string]string{"error": msj})
}

//GetID Obtiene un ID genérico
func GetID(w http.ResponseWriter, r *http.Request) {
	ID := bson.NewObjectId().Hex()
	RespondWithJSON(w, http.StatusOK, map[string]string{"ID": ID})
}
