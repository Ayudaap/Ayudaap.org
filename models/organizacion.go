package models

// Modelo de datos que representa una organizacion
type Organizacion struct {
	ID                 string           `json:"ID"`
	Tipo               TipoOrganizacion `json:"tipo,omitempty"`
	Nombre             string           `json:"nombre,omitempty"`
	RepresentanteLegal string           `json:"representanteLegal,omitempty"`
	Domicilio          Direccion        `json:"direccion,omitempty"`
}
