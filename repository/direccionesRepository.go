package repository

import (
	"Ayudaap.org/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GetDireccionByOrganizacionID Obtiene una Direccion por Id
func GetDireccionByOrganizacionID(id string) (models.Direccion, error) {
	col, ctx, cancel := GetCollection(organizacionCollection)
	defer cancel()

	Oid, _ := primitive.ObjectIDFromHex(id)

	var organizacion *models.Organizacion

	err := col.FindOne(ctx, bson.M{"_id": Oid}).Decode(&organizacion)
	if err != nil {
		return models.Direccion{}, err
	}

	return organizacion.Domicilio, nil
}

// UpdateDireccion Actualiza la direccion de una organizacion
func UpdateDireccion(id string, direccion *models.Direccion) (int64, error) {
	col, ctx, cancel := GetCollection(organizacionCollection)
	defer cancel()

	Oid, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": Oid}
	update := bson.M{"$set": bson.M{"direccion": direccion}}

	result, err := col.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}
