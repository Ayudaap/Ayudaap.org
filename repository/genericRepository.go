package repository

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Repositorio de base de datos
type GenericRepository struct {
	// Nombre de la colecion del modelo
	Collection string
}

// Inserta una nueva instancia del modelo
func (g *GenericRepository) Insert(modelo interface{}) string {
	col, ctx, cancel := GetCollection(DataBase, g.Collection)
	defer cancel()

	resultado, err := col.InsertOne(ctx, modelo)
	if err != nil {
		log.Fatal(err)
	}

	ObjectID, _ := resultado.InsertedID.(primitive.ObjectID)

	var result string = ObjectID.Hex()
	return result
}

// Obtiene todas las instancias del objeto
func (g *GenericRepository) GetAll() []interface{} {
	var modelos []interface{}

	col, ctx, cancel := GetCollection(DataBase, g.Collection)
	defer cancel()

	datos, err := col.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for datos.Next(ctx) {
		var modelo interface{}
		modelos = append(modelos, modelo)
	}
	return modelos
}

// Obtiene una objeto por Id
func (g *GenericRepository) GetModeloById(id string) *interface{} {
	col, ctx, cancel := GetCollection(DataBase, g.Collection)
	defer cancel()

	Oid, _ := primitive.ObjectIDFromHex(id)

	var modelo *interface{}
	err := col.FindOne(ctx, bson.M{"_id": Oid})
	if err != nil {
		return nil
	}

	return modelo
}

// Elimina una modelo
func (g *GenericRepository) Delete(id string) (int, error) {
	col, ctx, cancel := GetCollection(DataBase, g.Collection)
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

// Actualiza una modelo retornando el total de elementos que se modificaron
func (g *GenericRepository) Update(modelo *interface{}, id string) (int64, error) {

	col, ctx, cancel := GetCollection(DataBase, g.Collection)
	defer cancel()

	filter := bson.M{"_id": id}
	update := bson.M{"$set": modelo}

	result, err := col.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

//Purgar modelos borra toda la collecion
func (g *GenericRepository) Purgar() error {

	col, ctx, cancel := GetCollection(DataBase, g.Collection)
	defer cancel()

	err := col.Drop(ctx)

	return err
}
