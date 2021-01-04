package models

import (
	"time"

	"Ayudaap.org/pkg/database"
	"Ayudaap.org/src/entities"
)

//OrganizacionModel Tipo de organizacion
type OrganizacionModel struct{}

//ORGANIZACIONCOLLECTION Nombre de la conexion
const ORGANIZACIONCOLLECTION = "organizaciones"

var db database.GenericDB

func init() {
	db = database.GenericDB{CollectionName: ORGANIZACIONCOLLECTION}
}

//InsertOne Inserta un nuevo registro en la base de datos
func (o OrganizacionModel) InsertOne(organizacion entities.Organizacion) (string, error) {

	id, err := db.InsertOne(organizacion)
	if err != nil {
		return "", err
	}

	return id, nil
}

//FindAll Regresa todas las organizaciones
func (o OrganizacionModel) FindAll() ([]entities.Organizacion, error) {
	var resultados []entities.Organizacion
	err := db.FindAll(&resultados)

	if err != nil {
		return nil, err
	}

	return resultados, nil
}

//FindByID Encuentra una organizacion por ID
func (o OrganizacionModel) FindByID(ID string) (entities.Organizacion, error) {
	var organizacion entities.Organizacion
	err := db.FindByID(ID, &organizacion)

	if err != nil {
		return organizacion, err
	}

	return organizacion, nil
}

//Update Actualiza un objeto en la base de datos
func (o OrganizacionModel) Update(organizacion entities.Organizacion) (int64, error) {

	//TODO: Cambiar por codigo el usuario que lo modifica
	organizacion.Auditoria.ModificadoPor = "Testing"
	organizacion.Auditoria.UpdatedAt = time.Now().Unix()

	total, err := db.Udate(organizacion.ID.Hex(), organizacion)
	if err != nil {
		return 0, err
	}
	return total, nil

}

//DeleteOne Borra un proyecto por
func (o OrganizacionModel) DeleteOne(ID string) error {
	err := db.DeleteOne(ID)
	if err != nil {
		return err
	}
	return nil
}
