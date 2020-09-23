package test

import (
	"testing"

	repository "Ayudaap.org/Repository"
)

func proyectosTest(t *testing.T) {
	repositorio := repository.OrganizacionesRepository{DbRepo: *repository.GetInstance()}

	esperado := repositorio.GetAllOrganizaciones()
	if obtenido := len(esperado); obtenido <= 0 {
		t.Errorf("Esperado > 1, se obtuvo %d", obtenido)
	}

}
