package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Proyecto Modelo de datos que representa un proyecto
type Proyecto struct {
	ID        primitive.ObjectID `json:"Id" bson:"_id`
	Nombre    string             `json:"nombre" bson:"nombre"`
	Objetivo  string             `json:"objetivo,omitempty" bson:"objetivo"`
	Activo    bool               `json:"activo" bson:"activo"`
	Banner    string             `json:"banner" bson:"banner"`
	Area      Area               `json:"area" bson:"area"`
	Auditoria Auditoria          `json:"auditoria,omitempty" bson:"auditoria"`

	//TODO: Mover a modelo de Convocatoria
	Actividad             string  `json:"actividad,omitempty" bson:"actividad"`
	VoluntariosRequeridos int     `json:"voluntariosRequeridos" bson:"voluntariosRequeridos"`
	CapacidadesRequeridas string  `json:"capacidadesRequeridas,omitempty" bson:"capacidadesRequeridas"`
	Costo                 float32 `json:"costo,omitempty" bson:"costo"`
}
