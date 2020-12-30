package database

import (
	"testing"
	"time"

	"Ayudaap.org/pkg/config"
	"Ayudaap.org/src/entities"
)

func init() {
	config.Initconfig()
}

//Revisa la conexion a la base de datos que sea correcta
func TestConnection(t *testing.T) {

	_ = GetInstance()
	if i := ChequeoConnection(); i == 0 {
		t.Errorf("No se pudo conectar a la base de datos")
	}
}

//Revisa que se obtenga la coleccion correctamente
func TestGetCollection(t *testing.T) {

	esperado := "proyectos"
	col, _, cancel := GetCollection(esperado)
	defer cancel()

	recibido := col.Name()

	if recibido != esperado {
		t.Errorf("Esperado: %s | Recibido: %s", esperado, col.Name())
	}
}

func TestGetAuditoria(t *testing.T) {

	esperado := entities.Auditoria{
		CreatedAt:     time.Now().Unix(),
		UpdatedAt:     time.Now().Unix(),
		ModificadoPor: "erdvillegas",
	}

	recibido := GetAuditoria("erdvillegas")
	esperado.CreatedAt = recibido.CreatedAt
	esperado.UpdatedAt = recibido.UpdatedAt

	if esperado != recibido {
		t.Errorf("Esperado: %+v \nRecibido: %+v", esperado, recibido)
	}
}
