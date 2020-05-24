package routes

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"syreclabs.com/go/faker"
	"syreclabs.com/go/faker/locales"

	"Ayudaap.org/models"
	"Ayudaap.org/repository"
)

var organizaciones []models.Organizacion

// Inicializa la lista de organizaciones
func InicializarOrganizaciones(w http.ResponseWriter, r *http.Request) {
	faker.Locale = locales.En
	rand.Seed(50)
	total := rand.Intn(42)

	for i := 0; i <= total; i++ {

		organizaciones = append(organizaciones, models.Organizacion{
			ID:   primitive.NewObjectID(),
			Tipo: models.OrganizacionNoGubernamental,
			Domicilio: models.Direccion{
				ID:             primitive.NewObjectID(),
				Calle:          faker.Address().StreetName(),
				NumeroExterior: faker.Address().BuildingNumber(),
				CodigoPostal:   faker.Address().Postcode(),
				Colonia:        faker.Address().City(),
				Estado:         faker.Address().State(),
			},
			Nombre:             faker.Company().Name(),
			RepresentanteLegal: faker.Name().Name(),
		})

		rd := faker.RandomInt(1, 7)
		for j := 0; j <= rd; j++ {
			rndVal := faker.RandomInt(0, 1)
			var principal bool = false
			if rndVal == 1 {
				principal = true
			}

			organizaciones[i].Domicilio.Directorio = append(organizaciones[i].Domicilio.Directorio, models.Directorio{
				Alias:             fmt.Sprintf("%s %s", faker.Name().Prefix(), faker.Name().LastName()),
				CorreoElectronico: faker.Internet().Email(),
				Nombre:            faker.Name().Name(),
				Telefono:          faker.PhoneNumber().PhoneNumber(),
				EsPrincipal:       principal,
				ID:                primitive.NewObjectID(),
			})
		}
	}

	guardarOrganizacionesInicializer()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.RespuestaGenerica{fmt.Sprintf("Inizializado: %d objetos generados", total)})
}

// Inicializa la base de datos
func guardarOrganizacionesInicializer() {
	orgRepo := new(repository.OrganizacionesRepository)
	insertado := make(chan string)

	for _, org := range organizaciones {
		go orgRepo.InsertOrganizacion(org, insertado)
	}
}

// Lista todas las organizaciones
func GetALlOrganizacionesReq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(organizaciones)
	orgRepo := new(repository.OrganizacionesRepository)

	resultados := orgRepo.GetAllOrganizaciones()

	if len(resultados) <= 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.RespuestaGenerica{"No se encontraron datos a mostrar"})
	} else {
		json.NewEncoder(w).Encode(resultados)
	}
}

// Obtiene una organizacion por ID
func GetOrganizacionById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	orgRepo := new(repository.OrganizacionesRepository)
	resultados := orgRepo.GetOrganizacionById(id)

	if resultados == nil {
		json.NewEncoder(w).Encode(models.RespuestaGenerica{"No se encontraron datos a mostrar"})
	} else {
		json.NewEncoder(w).Encode(resultados)
	}
}
