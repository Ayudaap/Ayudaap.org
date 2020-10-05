package repository

import (
	"log"
	"time"

	"Ayudaap.org/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Nombre de la tabla de organizaciones
const organizacionCollection string = "organizaciones"

//InsertOrganizacion Inserta una nueva instancia de Organizacion
func InsertOrganizacion(organizacion models.Organizacion) string {
	col, ctx, cancel := GetCollection(organizacionCollection)
	defer cancel()

	organizacion.Auditoria = models.Auditoria{
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
func GetAllOrganizaciones() []models.Organizacion {
	var organizaciones []models.Organizacion

	col, ctx, cancel := GetCollection(organizacionCollection)
	defer cancel()

	datos, err := col.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for datos.Next(ctx) {
		var organizacion models.Organizacion
		err := datos.Decode(&organizacion)
		if err != nil {
			log.Fatal(err)
		}
		organizaciones = append(organizaciones, organizacion)
	}
	return organizaciones
}

//GetOrganizacionById Obtiene una organizacion por Id
func GetOrganizacionByID(id string) *models.Organizacion {
	col, ctx, cancel := GetCollection(organizacionCollection)
	defer cancel()

	Oid, _ := primitive.ObjectIDFromHex(id)

	var organizacion *models.Organizacion
	err := col.FindOne(ctx, bson.M{"_id": Oid}).Decode(&organizacion)
	if err != nil {
		return nil
	}

	return organizacion
}

//DeleteOrganizacion Elimina una organizacion
func DeleteOrganizacion(id string) (int, error) {
	col, ctx, cancel := GetCollection(organizacionCollection)
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

//UpdateOrganizacion Actualiza una organizacion retornando el total de elementos que se modificaron
func UpdateOrganizacion(organizacion *models.Organizacion) (int64, error) {

	col, ctx, cancel := GetCollection(organizacionCollection)
	defer cancel()

	filter := bson.M{"_id": organizacion.ID}
	update := bson.M{"$set": organizacion}

	organizacion.Auditoria.UpdatedAt = primitive.Timestamp{T: uint32(time.Now().Unix())}

	result, err := col.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

//PurgarOrganizaciones borra toda la collecion
func PurgarOrganizaciones() error {

	col, ctx, cancel := GetCollection(organizacionCollection)
	defer cancel()

	err := col.Drop(ctx)

	return err
}
