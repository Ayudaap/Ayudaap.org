package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	app "Ayudaap.org/common"
)

func main() {

	servidor := app.New()
	puerto := servidor.Puerto()
	r := servidor.Router()

	fmt.Printf("Ejecutando en :%d\n", puerto)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%d", puerto),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
