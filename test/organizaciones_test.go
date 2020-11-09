package test

import (
	"fmt"
	"testing"
	"time"

	"Ayudaap.org/src/entities"
	models "Ayudaap.org/src/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func proyectosTest(t *testing.T) {
	esperado, err := models.GetAllOrganizaciones()

	if err != nil {
		t.Errorf(err.Error())
	}

	if obtenido := len(esperado); obtenido <= 0 {
		t.Errorf("Esperado > 1, se obtuvo %d", obtenido)
	}
}

func organizacionesCrearTest(t *testing.T) {

	organizacion := entities.Organizacion{
		Auditoria: entities.Auditoria{
			CreatedAt:     primitive.Timestamp{T: uint32(time.Now().Unix())},
			UpdatedAt:     primitive.Timestamp{T: uint32(time.Now().Unix())},
			ModificadoPor: "1f96eb1d-bd3d-4c36-92b8-506c793a0731",
		},
		Banner: "https://picsum.photos/200/300",
		Nombre: "Agus Le Mechi",
		Tipo:   entities.OrganizacionSocialSinFinesDeLucro,
		Domicilio: entities.Direccion{
			Calle:          "Siempre Viva",
			CodigoPostal:   "66005",
			Colonia:        "Springfield",
			Estado:         "De prueba",
			NumeroExterior: "123",
			Referencia:     "Famosa casa animada",
			Directorio: []entities.Directorio{
				entities.Directorio{
					Alias:             "Hommer",
					CorreoElectronico: "homero.simpson@test.com",
					Nombre:            "Homero Simpson",
					Telefono:          "123-321-221",
					EsPrincipal:       true,
				},
			},
		},
	}

	repositorio, err := models.InsertOrganizacion(organizacion)

	fmt.Printf("Id: %s", repositorio)
	if err != nil {
		t.Errorf("No se pudo insertar el objeto")
	}

}
