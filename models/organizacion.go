package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Modelo de datos que representa una organizacion
type Organizacion struct {
	ID                 primitive.ObjectID `json:"Id" bson:"_id, omitempty"`
	Tipo               TipoOrganizacion   `json:"tipo,omitempty" bson:"tipo"`
	Nombre             string             `json:"nombre,omitempty" bson:"nombre"`
	RepresentanteLegal string             `json:"representanteLegal,omitempty" bson:"representanteLegal"`
	Domicilio          Direccion          `json:"direccion,omitempty" bson:"direccion"`
}
