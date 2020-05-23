package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"Ayudaap.org/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Repositorio de base de datos
type OrganizacionesRepository struct{}

// Nombre de la tabla de organizaciones
const orgCollection string = "organizaciones"

// Obtiene todas las organizaciones
func (o *OrganizacionesRepository) GetAllOrganizaciones() []models.Organizacion {
	var organizaciones []models.Organizacion

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	db := MongoCN.Database(DataBase)
	col := db.Collection(orgCollection)

	datos, err := col.Find(ctx, bson.D{})
	defer datos.Close(ctx)
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

func (o *OrganizacionesRepository) GetOrganizacionById(id string) *models.Organizacion {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(DataBase)
	col := db.Collection(orgCollection)

	_id, _ := primitive.ObjectIDFromHex(id)

	var organizacion *models.Organizacion
	err := col.FindOne(ctx, bson.M{"_id": _id}).Decode(&organizacion)
	if err != nil {
		return nil
	}

	return organizacion
}

// Inserta una nueva instancia de Organizacion
func (o *OrganizacionesRepository) InsertOrganizacion(organizacion models.Organizacion, c chan string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(DataBase)
	col := db.Collection(orgCollection)

	datos, err := col.InsertOne(ctx, organizacion)
	if err != nil {
		log.Fatal(err)
		c <- ""
	}
	var result string = fmt.Sprint(datos.InsertedID)
	c <- result
}
