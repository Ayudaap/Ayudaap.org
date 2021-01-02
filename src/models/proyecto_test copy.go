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

func TestInsertOneProyecto(t *testing.T) {

	proyectoEsperado := entities.Proyecto{
		ID:        primitive.NewObjectID(),
		Nombre:    faker.Company().Name(),
		Objetivo: faker.
		Banner:    faker.Avatar().String(),
		Auditoria: database.GetAuditoria(faker.Internet().UserName()),
	}

	rd := faker.RandomInt(1, 7)
	for j := 0; j <= rd; j++ {
		rndVal := faker.RandomInt(0, 1)
		var principal bool = false
		if rndVal == 1 {
			principal = true
		}

		proyectoEsperado.Domicilio.Directorio = append(proyectoEsperado.Domicilio.Directorio, entities.Directorio{
			Alias:             fmt.Sprintf("%s %s", faker.Name().Prefix(), faker.Name().LastName()),
			CorreoElectronico: faker.Internet().Email(),
			Nombre:            faker.Name().Name(),
			Telefono:          faker.PhoneNumber().PhoneNumber(),
			EsPrincipal:       principal,
			ID:                primitive.NewObjectID(),
		})
	}

	recibido, err := OrganizacionModel{}.InsertOne(proyectoEsperado)
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

	organizaciones, err := OrganizacionModel{}.FindAll()
	if err != nil {
		t.Skip()
	}

	want := organizaciones[0].ID.Hex()
	got, err := OrganizacionModel{}.FindByID(want)

	if err != nil {
		t.Errorf("No se pudo ejecutar la consulta \n :%s", err.Error())
	}

	if got.Nombre == "" {
		t.Errorf("No se logro ejecutar la consulta, want: %s got: %s", want, got.ID.Hex())
	}

}

func TestDeleteOrganizacion(t *testing.T) {
	organizaciones, err := OrganizacionModel{}.FindAll()
	if err != nil {
		t.Skip()
	}

	toDelete := organizaciones[0].ID.Hex()

	err = nil
	err = OrganizacionModel{}.DeleteOne(toDelete)
	if err != nil {
		t.Errorf("No se pudo borrar el registro:  %s", err.Error())
	}
}

func TestUpdateOrganizacion(t *testing.T) {
	organizaciones, err := OrganizacionModel{}.FindAll()
	if err != nil {
		t.Skip()
	}

	toUpdate := organizaciones[0]
	if toUpdate.Nombre == "" {
		t.Skip()
	}

	toUpdate.Nombre = "Prueba 1"
	toUpdate.Tipo = entities.OrganizacionNoGubernamental

	modificdo, err := OrganizacionModel{}.Update(&toUpdate)
	if err != nil {
		t.Skip()
	}

	if modificdo == 0 {
		t.Errorf("No se ejecuto modificacion")
	}
}

//TestUpdateOrganizacionDireccion Prueba modificar los datos de un documento envevido
func TestUpdateOrganizacionDireccion(t *testing.T) {
	organizaciones, err := OrganizacionModel{}.FindAll()
	if err != nil {
		t.Skip()
	}

	toUpdate := organizaciones[0]
	if toUpdate.Nombre == "" {
		t.Skip()
	}

	toUpdate.Domicilio.Calle = "Prueba Modificacion"
	toUpdate.Tipo = entities.OrganizacionNoGubernamental

	modificdo, err := OrganizacionModel{}.Update(&toUpdate)
	if err != nil {
		t.Skip()
	}

	if modificdo == 0 {
		t.Errorf("No se ejecuto modificacion")
	}
}

//TestUpdateOrganizacionDirectorio Prueba modificar los datos de un documento envevido
func TestUpdateOrganizacionDirectorio(t *testing.T) {
	organizaciones, err := OrganizacionModel{}.FindAll()
	if err != nil {
		t.Skip()
	}

	toUpdate := organizaciones[0]
	if toUpdate.Nombre == "" {
		t.Skip()
	}

	toUpdate.Domicilio.Directorio[0].Alias = "Prueba de modificacion"
	toUpdate.Tipo = entities.OrganizacionNoGubernamental

	modificdo, err := OrganizacionModel{}.Update(&toUpdate)
	if err != nil {
		t.Skip()
	}

	if modificdo == 0 {
		t.Errorf("No se ejecuto modificacion")
	}
}
