package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Contrato representa el documento
type Contrato struct {
	ID                 primitive.ObjectID `json:"Id" bson:"_id, omitempty"`
	FechaAceptado      time.Time          `json:"fechaAceptado" bson:"fechaAceptado"`
	InicioValidez      time.Time          `json:"inicioValidez" bson:"inicioValidez"`
	FinValidez         time.Time          `json:"finValidez" bson:"finValidez"`
	RepresentanteLegal string             `json:"representanteLegal,omitempty" bson:"representanteLegal"`
}
