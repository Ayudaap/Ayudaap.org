package models

import (
	"log"
	"time"

	"Ayudaap.org/src/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ProyectosCollection Nombre de la tabla de proyectos
const ProyectosCollection string = "proyectos"

//InsertProyecto Inserta una nueva instancia de Proyecto
func InsertProyecto(proyecto entities.Proyecto) string {
	col, ctx, cancel := GetCollection(ProyectosCollection)
	defer cancel()

	proyecto.Auditoria = entities.Auditoria{
		CreatedAt: primitive.Timestamp{T: uint32(time.Now().Unix())},
		UpdatedAt: primitive.Timestamp{T: uint32(time.Now().Unix())},
	}

	resultado, err := col.InsertOne(ctx, proyecto)
	if err != nil {
		log.Fatal(err)
	}

	ObjectID, _ := resultado.InsertedID.(primitive.ObjectID)

	var result string = ObjectID.Hex()
	return result
}

//GetAllProyectos Obtiene todas los proyectos
func GetAllProyectos() []entities.Proyecto {
	var proyectos []entities.Proyecto

	col, ctx, cancel := GetCollection(ProyectosCollection)
	defer cancel()

	datos, err := col.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for datos.Next(ctx) {
		var Proyecto entities.Proyecto
		err := datos.Decode(&Proyecto)
		if err != nil {
			log.Fatal(err)
		}
		proyectos = append(proyectos, Proyecto)
	}
	return proyectos
}

//GetProyectoByID Obtiene una Proyecto por Id
func GetProyectoByID(id string) *entities.Proyecto {
	col, ctx, cancel := GetCollection(ProyectosCollection)
	defer cancel()

	Oid, _ := primitive.ObjectIDFromHex(id)

	var Proyecto *entities.Proyecto
	err := col.FindOne(ctx, bson.M{"_Id": Oid}).Decode(&Proyecto)
	if err != nil {
		return nil
	}

	return Proyecto
}

//DeleteProyecto Elimina una Proyecto
func DeleteProyecto(id string) (int, error) {
	col, ctx, cancel := GetCollection(ProyectosCollection)
	defer cancel()

	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": oID}

	result, err := col.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	if result.DeletedCount <= 0 {
		return 0, nil
	}

	return int(result.DeletedCount), nil
}

//UpdateProyecto Actualiza una Proyecto retornando el total de elementos que se modificaron
func UpdateProyecto(proyecto *entities.Proyecto) (int64, error) {

	col, ctx, cancel := GetCollection(ProyectosCollection)
	defer cancel()

	filter := bson.M{"_id": proyecto.ID}
	update := bson.M{"$set": proyecto}

	proyecto.Auditoria.UpdatedAt = primitive.Timestamp{T: uint32(time.Now().Unix())}

	result, err := col.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

//PurgarProyectos Purgarproyectos borra toda la collecion
func PurgarProyectos() error {

	col, ctx, cancel := GetCollection(ProyectosCollection)
	defer cancel()

	err := col.Drop(ctx)

	return err
}
