package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Hola"))
}

var router = http.Handler

func New() *mux.Router{

	// TODO: Cambiar la forma de genear URL para version de API
	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/", home)

	// return router

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/",home)
	return router
}