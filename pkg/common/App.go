package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
}

// Implicacion de servidor
type server interface {
	Router() http.Handler
}

// Inicializa una nueva api
func New() server {
	a := &api{}

	r := mux.NewRouter()
	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}
