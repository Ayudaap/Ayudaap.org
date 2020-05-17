package app

import (
	"net/http"

	"Ayudaap.org/cmd/api/routes"
)

type api struct {
	router http.Handler
	puerto int
}

// Implicacion de servidor
type server interface {
	Router() http.Handler
	Puerto() int
}

// Inicializa una nueva api
func New() server {
	a := &api{}

	r := routes.GetHandler()
	a.router = r
	//TODO: Implementar extraccion por medio de archivo de configuracion
	a.puerto = 8081
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) Puerto() int {
	return a.puerto
}
