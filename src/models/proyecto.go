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

//PROYECTOS_proyectosCollectionCOLLECCION Nombre de la conexion
const proyectosCollection = "proyectos"

//InsertOne Inserta un nuevo registro en la base de datos
func (p ProyectoModel) InsertOne(proyecto entities.Proyecto) (string, error) {

	proyecto.ID = primitive.NewObjectID()
	proyecto.Area.ID = primitive.NewObjectID()
	if proyecto.Auditoria.CreatedAt == 0 {
		proyecto.Auditoria = database.GetAuditoria("NoName")
	}

	col, ctx, cancel := database.GetCollection(proyectosCollection)
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

	col, ctx, cancel := database.GetCollection(proyectosCollection)
	defer cancel()

	datos, err := col.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

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
	filter := bson.M{"_Id": oID}

	col, ctx, cancel := database.GetCollection(proyectosCollection)
	defer cancel()
	var proyecto entities.Proyecto

	err = col.FindOne(ctx, filter).Decode(&proyecto)

	return proyecto, err
}

//DeleteOne Borra un proyecto por
func (p ProyectoModel) DeleteOne(ID string) error {

	oID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_Id": oID}

	col, ctx, cancel := database.GetCollection(proyectosCollection)
	defer cancel()

	_, err := col.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}

//Purge Purga el proyecto
func (p ProyectoModel) Purge() error {
	col, ctx, cancel := database.GetCollection(proyectosCollection)
	defer cancel()

	err := col.Drop(ctx)
	if err != nil {
		return err
	} else {
		return nil
	}
}
