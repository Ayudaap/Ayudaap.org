package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

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
			ID:   faker.Bitcoin().Address(),
			Tipo: models.OrganizacionNoGubernamental,
			Domicilio: models.Direccion{
				ID:             faker.Bitcoin().Address(),
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
				ID:                faker.Bitcoin().Address(),
			})
		}
	}

	guardarOrganizacionesInicializer()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(struct {
		Mensaje string `json:"mensaje,omitempty"`
	}{"Ok"})
}

// Inicializa la base de datos
func guardarOrganizacionesInicializer() {
	log.Print("Inicializano base de datos")
	orgRepo := new(repository.OrganizacionesRepository)
	insertado := make(chan string)

	for _, org := range organizaciones {
		go orgRepo.InsertOrganizacion(org, insertado)
	}
}

// Lista todas las organizaciones
func GetALlOrganizacionesReq(w http.ResponseWriter, r *http.Request) {
	log.Printf("Peticion desde %s\n", r.RequestURI)
	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(organizaciones)
	orgRepo := new(repository.OrganizacionesRepository)
	resultados := orgRepo.GetAllOrganizaciones()

	json.NewEncoder(w).Encode(resultados)
}
