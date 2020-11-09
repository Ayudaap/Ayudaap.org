package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Auditoria del registro
type Auditoria struct {
	CreatedAt     primitive.Timestamp `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt     primitive.Timestamp `json:"updatedAt,omitempty" bson:"updatedAt"`
	ModificadoPor string              `json:"modificadoPor,omitempty" bson:"modificadoPor"`
}
