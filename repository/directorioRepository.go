package repository

import (
	"log"

	"Ayudaap.org/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//DirectorioRepository Repositorio de base de datos
type DirectorioRepository struct {
	DbRepo MongoRepository
}

//DirectorioCollection Nombre de la tabla de directorio
const DirectorioCollection string = "directorio"

//InsertarDirectorio Inserta una nueva instancia de Directorio
func (d *DirectorioRepository) InsertarDirectorio(directorio models.Directorio) string {
	col, ctx, cancel := d.DbRepo.GetCollection(DirectorioCollection)
	defer cancel()

	resultado, err := col.InsertOne(ctx, directorio)
	if err != nil {
		log.Fatal(err)
	}

	ObjectID, _ := resultado.InsertedID.(primitive.ObjectID)

	var result string = ObjectID.Hex()
	return result
}

//GetAllDirectorio Obtiene todas los directorio
func (d *DirectorioRepository) GetAllDirectorio() []models.Directorio {
	var directorio []models.Directorio

	col, ctx, cancel := d.DbRepo.GetCollection(DirectorioCollection)
	defer cancel()

	datos, err := col.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for datos.Next(ctx) {
		var Directorio models.Directorio
		err := datos.Decode(&Directorio)
		if err != nil {
			log.Fatal(err)
		}
		directorio = append(directorio, Directorio)
	}
	return directorio
}

//GetDirectorioByID Obtiene una Directorio por Id
func (d *DirectorioRepository) GetDirectorioByID(id string) *models.Directorio {
	col, ctx, cancel := d.DbRepo.GetCollection(DirectorioCollection)
	defer cancel()

	Oid, _ := primitive.ObjectIDFromHex(id)

	var Directorio *models.Directorio
	err := col.FindOne(ctx, bson.M{"_id": Oid}).Decode(&Directorio)
	if err != nil {
		return nil
	}

	return Directorio
}

//DeleteDirectorio Elimina una Directorio
func (d *DirectorioRepository) DeleteDirectorio(id string) (int, error) {
	col, ctx, cancel := d.DbRepo.GetCollection(DirectorioCollection)
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

//UpdateDirectorio Actualiza una Directorio retornando el total de elementos que se modificaron
func (d *DirectorioRepository) UpdateDirectorio(directorio *models.Directorio) (int64, error) {

	col, ctx, cancel := d.DbRepo.GetCollection(DirectorioCollection)
	defer cancel()

	filter := bson.M{"_id": directorio.ID}
	update := bson.M{"$set": directorio}

	result, err := col.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

//PurgarDirectorio PurgarDirectorio borra toda la collecion
func (d *DirectorioRepository) PurgarDirectorio() error {

	col, ctx, cancel := d.DbRepo.GetCollection(DirectorioCollection)
	defer cancel()

	err := col.Drop(ctx)

	return err
}
