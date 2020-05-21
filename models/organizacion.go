package models

// Modelo de datos que representa una organizacion
type Organizacion struct {
	ID                 string           `json:"ID" bson:"ID"`
	Tipo               TipoOrganizacion `json:"tipo,omitempty" bson:"tipo"`
	Nombre             string           `json:"nombre,omitempty" bson:"nombre"`
	RepresentanteLegal string           `json:"representanteLegal,omitempty" bson:"representanteLegal"`
	Domicilio          Direccion        `json:"direccion,omitempty" bson:"direccion"`
}
