package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"Ayudaap.org/models"
	"go.mongodb.org/mongo-driver/bson"
)

type OrganizacionesRepository struct{}

// Obtiene todas las organizaciones
func (o *OrganizacionesRepository) GetAllOrganizaciones() []models.Organizacion {
	var organizaciones []models.Organizacion

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	db := MongoCN.Database("tst")
	col := db.Collection("org")

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

// Inserta una nueva instancia de Organizacion
func (o *OrganizacionesRepository) InsertOrganizacion(organizacion models.Organizacion, c chan string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("tst")
	col := db.Collection("org")

	datos, err := col.InsertOne(ctx, organizacion)
	if err != nil {
		log.Fatal(err)
		c <- ""
	}
	var result string = fmt.Sprint(datos.InsertedID)
	log.Println(result)
	c <- result
}
