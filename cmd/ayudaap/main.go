package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	app "Ayudaap.org/cmd/common"
	"github.com/gorilla/handlers"
)

func main() {

	servidor := app.New()
	puerto := servidor.Puerto()
	r := servidor.Router()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})

	//Si esta en modo desarrollo se habilita para todos los origenes, en productivo se debe restringir
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	fmt.Printf("Ejecutando en :%d\n", puerto)

	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, handlers.CompressHandler(handlers.CORS(originsOk, headersOk, methodsOk)(r))),
		Addr:         fmt.Sprintf(":%d", puerto),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
