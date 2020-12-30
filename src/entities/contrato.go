package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Contrato representa el documento
type Contrato struct {
	ID                 primitive.ObjectID `json:"Id" bson:"_id, omitempty"`
	FechaAceptado      int64              `json:"fechaAceptado" bson:"fechaAceptado"`
	InicioValidez      int64              `json:"inicioValidez" bson:"inicioValidez"`
	FinValidez         int64              `json:"finValidez" bson:"finValidez"`
	RepresentanteLegal string             `json:"representanteLegal,omitempty" bson:"representanteLegal"`
}
