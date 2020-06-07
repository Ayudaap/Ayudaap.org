package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	app "Ayudaap.org/common"
	"github.com/gorilla/handlers"
)

func main() {

	servidor := app.New()
	puerto := servidor.Puerto()
	r := servidor.Router()

	fmt.Printf("Ejecutando en :%d\n", puerto)

	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, handlers.CompressHandler(r)),
		Addr:         fmt.Sprintf(":%d", puerto),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
