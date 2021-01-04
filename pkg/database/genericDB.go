package database

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"gopkg.in/mgo.v2/bson"
)

//GenericDB Colleccion generica para realizar las operaciones
type GenericDB struct {
	//Nombre de la coleccion
	CollectionName string
}

//InsertOne Inserta un nuevo registro en la base de datos
func (g GenericDB) InsertOne(registro interface{}) (string, error) {

	col, ctx, cancel := GetCollection(g.CollectionName)
	defer cancel()

	resultado, err := col.InsertOne(ctx, registro)
	if err != nil {
		return "", err
	}

	ObjectID, _ := resultado.InsertedID.(primitive.ObjectID)
	return ObjectID.Hex(), nil
}

//FindAll Regresa todos los registros
func (g GenericDB) FindAll(resultados interface{}) error {

	col, ctx, cancel := GetCollection(g.CollectionName)
	defer cancel()

	datos, err := col.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	err = datos.All(ctx, resultados)
	if err != nil {
		return err
	}

	return nil
}

//FindByID Encuentra una registro por ID
func (g GenericDB) FindByID(ID string, resultado interface{}) error {

	oID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err.Error())
	}

	col, ctx, cancel := GetCollection(g.CollectionName)
	defer cancel()

	err = col.FindOne(ctx, bson.M{"_id": oID}).Decode(resultado)
	if err != nil {
		log.Println(err.Error())
	}

	return err
}

//DeleteOne Borra un registro por ID
func (g GenericDB) DeleteOne(ID string) error {

	oID, _ := primitive.ObjectIDFromHex(ID)
	col, ctx, cancel := GetCollection(g.CollectionName)
	defer cancel()

	_, err := col.DeleteOne(ctx, bson.M{"_id": oID})

	if err != nil {
		return err
	}

	return nil
}

//Udate Actualiza un registro
func (g GenericDB) Udate(ID string, registro interface{}) (int64, error) {

	col, ctx, cancel := GetCollection(g.CollectionName)
	defer cancel()

	oID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}

	result, err := col.UpdateOne(ctx, bson.M{"_id": oID}, bson.M{"$set": registro})
	if err != nil {
		return 0, nil
	}

	return result.ModifiedCount, nil

}

//Purge Purga la coleccion de la base de datos
func (g GenericDB) Purge() error {
	col, ctx, cancel := GetCollection(g.CollectionName)
	defer cancel()

	err := col.Drop(ctx)
	if err != nil {
		return err
	}

	return nil
}
