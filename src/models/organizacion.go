package models

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"Ayudaap.org/pkg/database"
	"Ayudaap.org/src/entities"
	"gopkg.in/mgo.v2/bson"
)

//OrganizacionModel Tipo de organizacion
type OrganizacionModel struct{}

//organizacionCollection Nombre de la conexion
const organizacionCollection = "organizacion"

//InsertOne Inserta un nuevo registro en la base de datos
func (p OrganizacionModel) InsertOne(organizacion entities.Organizacion) (string, error) {

	organizacion.ID = primitive.NewObjectID()
	if organizacion.Auditoria.CreatedAt == 0 {
		organizacion.Auditoria = database.GetAuditoria("NoName")
	}

	col, ctx, cancel := database.GetCollection(organizacionCollection)
	defer cancel()

	resultado, err := col.InsertOne(ctx, organizacion)
	if err != nil {
		return "", err
	}

	ObjectID, _ := resultado.InsertedID.(primitive.ObjectID)
	return ObjectID.Hex(), nil
}

//FindAll Regresa todas las organizaciones
func (p OrganizacionModel) FindAll() ([]entities.Organizacion, error) {
	var organizaciones []entities.Organizacion

	col, ctx, cancel := database.GetCollection(organizacionCollection)
	defer cancel()

	datos, err := col.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	for datos.Next(ctx) {
		var organizacion entities.Organizacion
		err := datos.Decode(&organizacion)
		if err != nil {
			log.Fatal(err)
		}
		organizaciones = append(organizaciones, organizacion)
	}

	return organizaciones, nil
}
