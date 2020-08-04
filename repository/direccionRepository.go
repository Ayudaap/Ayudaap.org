package repository

import (
	"log"

	"Ayudaap.org/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//DireccionesRepository Repositorio de base de datos
type DireccionesRepository struct {
	DbRepo MongoRepository
}

//DireccionesCollection Nombre de la tabla de Direcciones
const DireccionesCollection string = "direcciones"

//InsertDireccion Inserta una nueva instancia de Direccion
func (o *DireccionesRepository) InsertDireccion(Direccion models.Direccion) string {
	col, ctx, cancel := o.DbRepo.GetCollection(DireccionesCollection)
	defer cancel()

	resultado, err := col.InsertOne(ctx, Direccion)
	if err != nil {
		log.Fatal(err)
	}

	ObjectID, _ := resultado.InsertedID.(primitive.ObjectID)

	var result string = ObjectID.Hex()
	return result
}

//GetAllDirecciones Obtiene todas los Direcciones
func (o *DireccionesRepository) GetAllDirecciones() []models.Direccion {
	var Direcciones []models.Direccion

	col, ctx, cancel := o.DbRepo.GetCollection(DireccionesCollection)
	defer cancel()

	datos, err := col.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for datos.Next(ctx) {
		var Direccion models.Direccion
		err := datos.Decode(&Direccion)
		if err != nil {
			log.Fatal(err)
		}
		Direcciones = append(Direcciones, Direccion)
	}
	return Direcciones
}

//GetDireccionByID Obtiene una Direccion por Id
func (o *DireccionesRepository) GetDireccionByID(id string) *models.Direccion {
	col, ctx, cancel := o.DbRepo.GetCollection(DireccionesCollection)
	defer cancel()

	Oid, _ := primitive.ObjectIDFromHex(id)

	var Direccion *models.Direccion
	err := col.FindOne(ctx, bson.M{"_id": Oid}).Decode(&Direccion)
	if err != nil {
		return nil
	}

	return Direccion
}

//DeleteDireccion Elimina una Direccion
func (o *DireccionesRepository) DeleteDireccion(id string) (int, error) {
	col, ctx, cancel := o.DbRepo.GetCollection(DireccionesCollection)
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

//UpdateDireccion Actualiza una Direccion retornando el total de elementos que se modificaron
func (o *DireccionesRepository) UpdateDireccion(Direccion *models.Direccion) (int64, error) {

	col, ctx, cancel := o.DbRepo.GetCollection(DireccionesCollection)
	defer cancel()

	filter := bson.M{"_id": Direccion.ID}
	update := bson.M{"$set": Direccion}

	result, err := col.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

//PurgarDirecciones PurgarDirecciones borra toda la collecion
func (o *DireccionesRepository) PurgarDirecciones() error {

	col, ctx, cancel := o.DbRepo.GetCollection(DireccionesCollection)
	defer cancel()

	err := col.Drop(ctx)

	return err
}
