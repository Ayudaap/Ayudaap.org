package api

import (
	"encoding/json"
	"net/http"
)

//respondWithError Responde con un mensaje de error
func respondWithError(w http.ResponseWriter, code int, msj string) {
	respondWithJSON(w, code, map[string]string{"error": msj})
}

//respondWithJSON Responde con un mensaje de error
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	// response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	// w.Write(response)
	json.NewEncoder(w).Encode(payload)
}
