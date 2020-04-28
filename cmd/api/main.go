package main

import (
	"fmt"
	"net/http"

	app "Ayudaap.org/pkg/common"
)

func main() {

	puerto := 8081
	servidor := app.New()
	fmt.Printf("Ejecutando en :%d\n", puerto)
	http.ListenAndServe(fmt.Sprintf(":%d", puerto), servidor.Router())
}
