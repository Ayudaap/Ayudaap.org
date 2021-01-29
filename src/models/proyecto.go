package models

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"Ayudaap.org/pkg/database"
	"Ayudaap.org/src/entities"
	"gopkg.in/mgo.v2/bson"
)

//ProyectoModel Tipo de proyecto
type ProyectoModel struct{}

//PROYECTOSCOLLECTION Nombre de la conexion
const PROYECTOSCOLLECTION = "proyectos"

//InsertOne Inserta un nuevo registro en la base de datos
func (p ProyectoModel) InsertOne(proyecto entities.Proyecto) (string, error) {

	proyecto.ID = primitive.NewObjectID()
	proyecto.Area.ID = primitive.NewObjectID()
	if proyecto.Auditoria.CreatedAt == 0 {
		proyecto.Auditoria = database.GetAuditoria("NoName")
	}

	col, ctx, cancel := database.GetCollection(PROYECTOSCOLLECTION)
	defer cancel()

	resultado, err := col.InsertOne(ctx, proyecto)
	if err != nil {
		return "", err
	}

	ObjectID, _ := resultado.InsertedID.(primitive.ObjectID)
	return ObjectID.Hex(), nil
}

//FindAll Regresa todos los proyectos
func (p ProyectoModel) FindAll() ([]entities.Proyecto, error) {
	var proyectos []entities.Proyecto

	col, ctx, cancel := database.GetCollection(PROYECTOSCOLLECTION)
	defer cancel()

	datos, err := col.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	defer datos.Close(ctx)
	for datos.Next(ctx) {
		var proyecto entities.Proyecto
		err := datos.Decode(&proyecto)
		if err != nil {
			log.Fatal(err)
		}
		proyectos = append(proyectos, proyecto)
	}

	return proyectos, nil
}

//FindByID Encuentra una proyecto por ID
func (p ProyectoModel) FindByID(ID string) (entities.Proyecto, error) {

	oID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err.Error())
	}

	col, ctx, cancel := database.GetCollection(PROYECTOSCOLLECTION)
	defer cancel()
	var proyecto entities.Proyecto

	err = col.FindOne(ctx, bson.M{"_Id": oID}).Decode(&proyecto)

	return proyecto, err
}

//DeleteOne Borra un proyecto por ID
func (p ProyectoModel) DeleteOne(ID string) error {

	oID, _ := primitive.ObjectIDFromHex(ID)
	col, ctx, cancel := database.GetCollection(PROYECTOSCOLLECTION)
	defer cancel()

	_, err := col.DeleteOne(ctx, bson.M{"_Id": oID})

	if err != nil {
		return err
	}

	return nil
}

//Update Actualiza un objeto
func (p ProyectoModel) Update(proyecto entities.Proyecto) (int64, error) {

	col, ctx, cancel := database.GetCollection(PROYECTOSCOLLECTION)
	defer cancel()

	result, err := col.UpdateOne(ctx, bson.M{"_Id": proyecto.ID}, bson.M{"$set": proyecto})
	if err != nil {
		return 0, nil
	}

	return result.ModifiedCount, nil

}
