package repository

import (
	"log"

	"Ayudaap.org/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Repositorio de base de datos
type OrganizacionesRepository struct{}

// Nombre de la tabla de organizaciones
const organizacionCollection string = "organizaciones"

// Inserta una nueva instancia de Organizacion
func (o *OrganizacionesRepository) InsertOrganizacion(organizacion models.Organizacion, c chan string) {
	col, ctx, cancel := GetCollection(DataBase, organizacionCollection)
	defer cancel()

	resultado, err := col.InsertOne(ctx, organizacion)
	if err != nil {
		log.Fatal(err)
	}

	ObjectID, _ := resultado.InsertedID.(primitive.ObjectID)

	var result string = ObjectID.Hex()
	c <- result
}

// Inserta una nueva instancia de Organizacion
// `Organizaciones` Arreglo de organizaciones que se insertaran
func (o *OrganizacionesRepository) InsertOrganizaciones(organizaciones []interface{}, c chan bool) {
	col, ctx, cancel := GetCollection(DataBase, organizacionCollection)
	defer cancel()

	resultado, err := col.InsertMany(ctx, organizaciones)
	if err != nil {
		log.Fatal(err)
	}

	c <- len(resultado.InsertedIDs) > 0
}

// Obtiene todas las organizaciones
func (o *OrganizacionesRepository) GetAllOrganizaciones() []models.Organizacion {
	var organizaciones []models.Organizacion

	col, ctx, cancel := GetCollection(DataBase, organizacionCollection)
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

// Obtiene una organizacion por Id
func (o *OrganizacionesRepository) GetOrganizacionById(id string) *models.Organizacion {
	col, ctx, cancel := GetCollection(DataBase, organizacionCollection)
	defer cancel()

	Oid, _ := primitive.ObjectIDFromHex(id)

	var organizacion *models.Organizacion
	err := col.FindOne(ctx, bson.M{"_id": Oid}).Decode(&organizacion)
	if err != nil {
		return nil
	}

	return organizacion
}

// Elimina una organizacion
func (o *OrganizacionesRepository) DeleteOrganizacion(id string) (int, error) {
	col, ctx, cancel := GetCollection(DataBase, organizacionCollection)
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

// Actualiza una organizacion retornando el total de elementos que se modificaron
func (o *OrganizacionesRepository) UpdateOrganizacion(organizacion *models.Organizacion) (int64, error) {

	col, ctx, cancel := GetCollection(DataBase, organizacionCollection)
	defer cancel()

	filter := bson.M{"_id": organizacion.ID}
	update := bson.M{"$set": organizacion}

	result, err := col.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}
