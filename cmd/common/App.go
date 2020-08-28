package app

import (
	"net/http"

	routes "Ayudaap.org/cmd/rutas"
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

	r := routes.GetRutas()
	a.router = r
	//TODO: Implementar extraccion por medio de archivo de configuracion
	a.puerto = 8080
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) Puerto() int {
	return a.puerto
}
