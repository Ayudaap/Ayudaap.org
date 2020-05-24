package routes

import (
	"encoding/json"
	"errors"
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

// Modelo de organizaciones
var orgRepo *repository.OrganizacionesRepository

func init() {
	orgRepo = new(repository.OrganizacionesRepository)
}

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
	resultados := orgRepo.GetOrganizacionById(id)

	if resultados == nil {
		json.NewEncoder(w).Encode(models.RespuestaGenerica{"No se encontraron datos a mostrar"})
	} else {
		json.NewEncoder(w).Encode(resultados)
	}
}

// Crea una nueva organizacion
func CreateOrganizacion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	var organizacion models.Organizacion

	if err := json.NewDecoder(r.Body).Decode(&organizacion); err != nil {
		GetError(err, w)
	}

	organizacion.ID = primitive.NewObjectID()
	organizacion.Domicilio.ID = primitive.NewObjectID()
	for _, dir := range organizacion.Domicilio.Directorio {
		dir.ID = primitive.NewObjectID()
	}

	organizacionInsertada := make(chan string)

	go orgRepo.InsertOrganizacion(organizacion, organizacionInsertada)
	idInsertado := <-organizacionInsertada

	if len(idInsertado) <= 0 {
		err := errors.New("No se pudo insertar el objeto")
		GetError(err, w)
	} else {
		json.NewEncoder(w).Encode(struct {
			Id string `json:"id,omitempty"`
		}{idInsertado})
	}

}
