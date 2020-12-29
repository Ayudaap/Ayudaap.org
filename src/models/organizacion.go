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

//ORGANIZACIONCOLLECTION Nombre de la conexion
const ORGANIZACIONCOLLECTION = "organizaciones"

//InsertOne Inserta un nuevo registro en la base de datos
func (o OrganizacionModel) InsertOne(organizacion entities.Organizacion) (string, error) {

	if organizacion.ID.IsZero() {
		organizacion.ID = primitive.NewObjectID()
	}

	if organizacion.Auditoria.CreatedAt == 0 {
		organizacion.Auditoria = database.GetAuditoria("NoName")
	}

	col, ctx, cancel := database.GetCollection(ORGANIZACIONCOLLECTION)
	defer cancel()

	resultado, err := col.InsertOne(ctx, organizacion)
	if err != nil {
		return "", err
	}

	ObjectID, _ := resultado.InsertedID.(primitive.ObjectID)
	return ObjectID.Hex(), nil
}

//FindAll Regresa todas las organizaciones
func (o OrganizacionModel) FindAll() ([]entities.Organizacion, error) {
	var organizaciones []entities.Organizacion

	col, ctx, cancel := database.GetCollection(ORGANIZACIONCOLLECTION)
	defer cancel()

	datos, err := col.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	defer datos.Close(ctx)

	err = datos.All(ctx, &organizaciones)
	if err != nil {
		return nil, err
	}

	return organizaciones, nil
}

//FindByID Encuentra una organizacion por ID
func (o OrganizacionModel) FindByID(id primitive.ObjectID) (entities.Organizacion, error) {
	filter := bson.M{"_id": id.Hex()}

	col, ctx, cancel := database.GetCollection(ORGANIZACIONCOLLECTION)
	defer cancel()
	var organizacion entities.Organizacion

	err := col.FindOne(ctx, filter).Decode(&organizacion)

	return organizacion, err
}

//DeleteOne Borra un proyecto por
func (o OrganizacionModel) DeleteOne(id primitive.ObjectID) error {

	col, ctx, cancel := database.GetCollection(ORGANIZACIONCOLLECTION)
	defer cancel()

	_, err := col.DeleteOne(ctx, bson.M{"_id": id.Hex()})

	if err != nil {
		return err
	}

	return nil
}

// func (o OrganizacionModel) GetOrganizacionByNombre(nombre string) (entities.Organizacion, error) {
// 	col, ctx, cancel := database.GetCollection(ORGANIZACIONCOLLECTION)
// 	defer cancel()

// 	_, err := col.Find
// }
