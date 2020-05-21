package repository

import (
	"context"
	"log"
	"time"

	"Ayudaap.org/models"
	"go.mongodb.org/mongo-driver/bson"
)

type OrganizacionesRepository struct{}

// Obtiene todas las organizaciones
func (o *OrganizacionesRepository) GetAllOrganizaciones() []models.Organizacion {

	_ = o
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
