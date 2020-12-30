package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Organizacion Modelo de datos que representa una organizacion
type Organizacion struct {
	ID        primitive.ObjectID `json:"Id" bson:"_id, omitempty"`
	Tipo      TipoOrganizacion   `json:"tipo,omitempty" bson:"tipo"`
	Nombre    string             `json:"nombre,omitempty" bson:"nombre"`
	Domicilio Direccion          `json:"direccion,omitempty" bson:"direccion"`
	Auditoria Auditoria          `json:"auditoria,omitempty" bson:"auditoria"`
	Banner    string             `json:"banner,omitempty" bson:"banner"`
}
