package organizaciones

import (
	"log"
	"time"

	au "Ayudaap.org/pkg/auditoria"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertOrganizacion Inserta una nueva instancia de Organizacion
func InsertOrganizacion(organizacion Organizacion) string {

	db := getConexion()
	col, ctx, cancel := db.GetCollection(organizacionCollection)
	defer cancel()

	organizacion.Auditoria = au.Auditoria{
		CreatedAt: primitive.Timestamp{T: uint32(time.Now().Unix())},
		UpdatedAt: primitive.Timestamp{T: uint32(time.Now().Unix())},
	}

	resultado, err := col.InsertOne(ctx, organizacion)
	if err != nil {
		log.Fatal(err)
	}

	ObjectID, _ := resultado.InsertedID.(primitive.ObjectID)

	var result string = ObjectID.Hex()
	return result
}

//GetAllOrganizaciones Obtiene todas las organizaciones
func GetAllOrganizaciones() []Organizacion {
	var organizaciones []Organizacion

	db := getConexion()
	col, ctx, cancel := db.GetCollection(organizacionCollection)
	defer cancel()

	datos, err := col.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for datos.Next(ctx) {
		var organizacion Organizacion
		err := datos.Decode(&organizacion)
		if err != nil {
			log.Fatal(err)
		}
		organizaciones = append(organizaciones, organizacion)
	}
	return organizaciones
}
