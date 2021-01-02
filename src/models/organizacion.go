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
		log.Println(err.Error())
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
func (o OrganizacionModel) FindByID(ID string) (entities.Organizacion, error) {
	oID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Println(err.Error())
		return entities.Organizacion{}, err
	}

	col, ctx, cancel := database.GetCollection(ORGANIZACIONCOLLECTION)
	defer cancel()
	var organizacion entities.Organizacion

	err = col.FindOne(ctx, bson.M{"_id": oID}).Decode(&organizacion)

	return organizacion, err
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
