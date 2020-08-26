package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Modelo de datos que representa una area
type Area struct {
	ID          primitive.ObjectID `json:"Id" bson:"_id, omitempty"`
	Nombre      string             `json:"nombre,omitempty" bson:"nombre"`
	Descripcion string             `json:"descripcion,omitempty" bson:"descripcion"`
	Auditoria   Auditoria          `json:"auditoria,omitempty" bson:"auditoria"`
}
