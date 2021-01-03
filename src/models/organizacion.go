package models

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"Ayudaap.org/pkg/database"
	"Ayudaap.org/src/entities"
	"gopkg.in/mgo.v2/bson"
)

//OrganizacionModel Tipo de organizacion
type OrganizacionModel struct{}

//ORGANIZACIONCOLLECTION Nombre de la conexion
const ORGANIZACIONCOLLECTION = "organizaciones"

var db database.GenericDB

func init() {
	db = database.GenericDB{CollectionName: ORGANIZACIONCOLLECTION}
}

//InsertOne Inserta un nuevo registro en la base de datos
func (o OrganizacionModel) InsertOne(organizacion entities.Organizacion) (string, error) {

	id, err := db.InsertOne(organizacion)
	if err != nil {
		return "", err
	}

	return id, nil
}

//FindAll Regresa todas las organizaciones
func (o OrganizacionModel) FindAll() ([]entities.Organizacion, error) {
	var resultados []entities.Organizacion
	err := db.FindAll(&resultados)

	if err != nil {
		return nil, err
	}

	return resultados, nil
}

//FindByID Encuentra una organizacion por ID
func (o OrganizacionModel) FindByID(ID string) (entities.Organizacion, error) {
	var organizacion entities.Organizacion
	err := db.FindByID(ID, &organizacion)

	if err != nil {
		return organizacion, err
	}

	return organizacion, nil
}

//Update Actualiza un objeto en la base de datos
func (o OrganizacionModel) Update(organizacion *entities.Organizacion) (int64, error) {

	col, ctx, cancel := database.GetCollection(ORGANIZACIONCOLLECTION)
	defer cancel()

	//TODO: Cambiar por codigo el usuario que lo modifica
	organizacion.Auditoria.ModificadoPor = "Testing"
	organizacion.Auditoria.UpdatedAt = time.Now().Unix()

	res, err := col.UpdateOne(ctx, bson.M{"_id": organizacion.ID}, bson.M{"$set": organizacion})

	if err != nil {
		return 0, err
	}

	return res.MatchedCount, nil

}

//DeleteOne Borra un proyecto por
func (o OrganizacionModel) DeleteOne(ID string) error {
	oID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Println(err.Error())
	}

	col, ctx, cancel := database.GetCollection(ORGANIZACIONCOLLECTION)
	defer cancel()

	_, err = col.DeleteOne(ctx, bson.M{"_id": oID})

	if err != nil {
		return err
	}

	return nil
}

//getID Convierte un string en ObjectID
func getID(ID string) (primitive.ObjectID, error) {
	oID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return oID, nil
}
