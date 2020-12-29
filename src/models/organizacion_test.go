package models

import (
	"fmt"
	"math/rand"
	"testing"

	"Ayudaap.org/src/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"syreclabs.com/go/faker"
	"syreclabs.com/go/faker/locales"

	"Ayudaap.org/pkg/config"
	"Ayudaap.org/pkg/database"
)

func init() {
	config.Initconfig()
	faker.Locale = locales.Es
	rand.Seed(50)
}

func TestInsertOneOrganizacion(t *testing.T) {

	id, _ := primitive.ObjectIDFromHex("5fe9889d8dbcbd4f4e5b5f45")

	organizacionEsperado := entities.Organizacion{
		ID:     id,
		Nombre: faker.Company().Name(),
		Banner: faker.Avatar().String(),
		Tipo:   entities.TipoOrganizacion(faker.RandomInt(0, 4)),
		Domicilio: entities.Direccion{
			ID:             primitive.NewObjectID(),
			Calle:          faker.Address().StreetName(),
			NumeroExterior: faker.Address().BuildingNumber(),
			CodigoPostal:   faker.Address().Postcode(),
			Colonia:        faker.Address().City(),
			Estado:         faker.Address().State(),
		},
		Auditoria: database.GetAuditoria(faker.Internet().UserName()),
	}

	rd := faker.RandomInt(1, 7)
	for j := 0; j <= rd; j++ {
		rndVal := faker.RandomInt(0, 1)
		var principal bool = false
		if rndVal == 1 {
			principal = true
		}

		organizacionEsperado.Domicilio.Directorio = append(organizacionEsperado.Domicilio.Directorio, entities.Directorio{
			Alias:             fmt.Sprintf("%s %s", faker.Name().Prefix(), faker.Name().LastName()),
			CorreoElectronico: faker.Internet().Email(),
			Nombre:            faker.Name().Name(),
			Telefono:          faker.PhoneNumber().PhoneNumber(),
			EsPrincipal:       principal,
			ID:                primitive.NewObjectID(),
		})
	}

	recibido, err := OrganizacionModel{}.InsertOne(organizacionEsperado)
	if err != nil {
		t.Errorf("No se pudo insertar el registro \n%s", err.Error())
	} else if "" == recibido {
		t.Error("Se recibio un id vacio")
	}
}

func TestGetAllOrganizacion(t *testing.T) {

	got, err := OrganizacionModel{}.FindAll()
	if err != nil {
		t.Errorf("No se pudo ejecutar la consulta")
	}

	if tam := len(got); tam == 0 {
		t.Errorf("No se recibieron datos de la consulta, esperado: >1 , recibido: %d", tam)
	}
}

func TestGetOneOrganizacion(t *testing.T) {

	want, _ := primitive.ObjectIDFromHex("5fe9889d8dbcbd4f4e5b5f45")
	got, err := OrganizacionModel{}.FindByID(want)

	if err != nil {
		t.Errorf("No se pudo ejecutar la consulta \n :%s", err.Error())
	}

	if got.Nombre == "" {
		t.Errorf("No se logro ejecutar la consulta, want: %s got: %s", want.Hex(), got.ID.Hex())
	}

}
