package models

import (
	"math/rand"
	"testing"
	"time"

	"Ayudaap.org/pkg/config"
	"Ayudaap.org/pkg/database"
	"Ayudaap.org/src/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"syreclabs.com/go/faker"
	"syreclabs.com/go/faker/locales"
)

func init() {
	config.Initconfig()
	faker.Locale = locales.Es
	rand.Seed(50)
}

func TestInsertOneProyecto(t *testing.T) {

	inicioDuracion, _ := time.ParseDuration("720h")
	finDuracion, _ := time.ParseDuration("1080h")
	var activo bool = false
	rndNum := faker.RandomInt(0, 1)
	if rndNum == 1 {
		activo = true
	}

	proyectoEsperado := entities.Proyecto{
		ID:                    primitive.NewObjectID(),
		Nombre:                faker.Name().Name(),
		Objetivo:              faker.Company().CatchPhrase(),
		Banner:                faker.Avatar().String(),
		Actividad:             faker.Hacker().SaySomethingSmart(),
		VoluntariosRequeridos: faker.Number().NumberInt(3),
		CapacidadesRequeridas: faker.Hacker().Phrases()[1],
		Costo:                 faker.Commerce().Price(),
		Inicio:                faker.Time().Forward(inicioDuracion).Unix(),
		Fin:                   faker.Time().Forward(finDuracion).Unix(),
		Activo:                activo,
		Area: entities.Area{
			ID:          primitive.NewObjectID(),
			Nombre:      faker.Commerce().Department(),
			Descripcion: faker.Commerce().ProductName(),
			Auditoria:   database.GetAuditoria(faker.Internet().UserName()),
		},
		Auditoria: database.GetAuditoria(faker.Internet().UserName()),
	}

	recibido, err := ProyectoModel{}.InsertOne(proyectoEsperado)
	if err != nil {
		t.Errorf("No se pudo insertar el registro \n%s", err.Error())
	} else if "" == recibido {
		t.Error("Se recibio un id vacio")
	}
}

func TestGetAllProyecto(t *testing.T) {

	got, err := ProyectoModel{}.FindAll()
	if err != nil {
		t.Errorf("No se pudo ejecutar la consulta")
	}

	if tam := len(got); tam == 0 {
		t.Errorf("No se recibieron datos de la consulta, esperado: >1 , recibido: %d", tam)
	}
}

// func TestGetOneOrganizacion(t *testing.T) {

// 	organizaciones, err := OrganizacionModel{}.FindAll()
// 	if err != nil {
// 		t.Skip()
// 	}

// 	want := organizaciones[0].ID.Hex()
// 	got, err := OrganizacionModel{}.FindByID(want)

// 	if err != nil {
// 		t.Errorf("No se pudo ejecutar la consulta \n :%s", err.Error())
// 	}

// 	if got.Nombre == "" {
// 		t.Errorf("No se logro ejecutar la consulta, want: %s got: %s", want, got.ID.Hex())
// 	}

// }

func TestDeleteProyecto(t *testing.T) {
	proyectos, err := ProyectoModel{}.FindAll()
	if err != nil {
		t.Skip()
	}

	toDelete := proyectos[0].ID.Hex()

	err = nil
	err = ProyectoModel{}.DeleteOne(toDelete)
	if err != nil {
		t.Errorf("No se pudo borrar el registro:  %s", err.Error())
	}
}

// func TestUpdateOrganizacion(t *testing.T) {
// 	organizaciones, err := OrganizacionModel{}.FindAll()
// 	if err != nil {
// 		t.Skip()
// 	}

// 	toUpdate := organizaciones[0]
// 	if toUpdate.Nombre == "" {
// 		t.Skip()
// 	}

// 	toUpdate.Nombre = "Prueba 1"
// 	toUpdate.Tipo = entities.OrganizacionNoGubernamental

// 	modificdo, err := OrganizacionModel{}.Update(toUpdate)
// 	if err != nil {
// 		t.Skip()
// 	}

// 	if modificdo == 0 {
// 		t.Errorf("No se ejecuto modificacion")
// 	}
// }

// //TestUpdateOrganizacionDireccion Prueba modificar los datos de un documento envevido
// func TestUpdateOrganizacionDireccion(t *testing.T) {
// 	organizaciones, err := OrganizacionModel{}.FindAll()
// 	if err != nil {
// 		t.Skip()
// 	}

// 	toUpdate := organizaciones[0]
// 	if toUpdate.Nombre == "" {
// 		t.Skip()
// 	}

// 	toUpdate.Domicilio.Calle = "Prueba Modificacion"
// 	toUpdate.Tipo = entities.OrganizacionNoGubernamental

// 	modificdo, err := OrganizacionModel{}.Update(toUpdate)
// 	if err != nil {
// 		t.Skip()
// 	}

// 	if modificdo == 0 {
// 		t.Errorf("No se ejecuto modificacion")
// 	}
// }

// //TestUpdateOrganizacionDirectorio Prueba modificar los datos de un documento envevido
// func TestUpdateOrganizacionDirectorio(t *testing.T) {
// 	organizaciones, err := OrganizacionModel{}.FindAll()
// 	if err != nil {
// 		t.Skip()
// 	}

// 	toUpdate := organizaciones[0]
// 	if toUpdate.Nombre == "" {
// 		t.Skip()
// 	}

// 	toUpdate.Domicilio.Directorio[0].Alias = "Prueba de modificacion"
// 	toUpdate.Tipo = entities.OrganizacionNoGubernamental

// 	modificdo, err := OrganizacionModel{}.Update(toUpdate)
// 	if err != nil {
// 		t.Skip()
// 	}

// 	if modificdo == 0 {
// 		t.Errorf("No se ejecuto modificacion")
// 	}
// }
