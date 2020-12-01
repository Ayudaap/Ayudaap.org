package entities

// Auditoria del registro
type Auditoria struct {
	CreatedAt     int64  `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt     int64  `json:"updatedAt,omitempty" bson:"updatedAt"`
	ModificadoPor string `json:"modificadoPor,omitempty" bson:"modificadoPor"`
}
