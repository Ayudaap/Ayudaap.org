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
