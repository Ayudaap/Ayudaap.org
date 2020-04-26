package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	log.Printf("Nueva peticion IP = %v", r.RemoteAddr)
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	mensaje := os.Args[0]
	fmt.Printf("Ejecutando desde :%s\n", mensaje)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	http.ListenAndServe(":8080", router)
}
