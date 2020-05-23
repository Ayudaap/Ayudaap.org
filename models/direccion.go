package models

// Directorio de contacto
type Directorio struct {
	ID                string `json:"ID" bson:"ID"`
	Nombre            string `json:"nombre" bson:"nombre"`
	Telefono          string `json:"telefono" bson:"telefono"`
	EsPrincipal       bool   `json:"esPrincipal" bson:"esPrincipal"`
	CorreoElectronico string `json:"correoElectronico" bson:"correoElectronico"`
	Alias             string `json:"alias" bson:"alias"`
}

// Direccion de la oficina
type Direccion struct {
	ID             string       `json:"ID" bson:"ID"`
	Calle          string       `json:"calle,omitempty" bson:"calle"`
	NumeroInterior string       `json:"numeroInterior,omitempty" bson:"numeroInterior"`
	NumeroExterior string       `json:"numeroExterior,omitempty" bson:"numeroExterior"`
	Colonia        string       `json:"colonia,omitempty" bson:"colonia"`
	CodigoPostal   string       `json:"codigoPostal,omitempty" bson:"codigoPostal"`
	Estado         string       `json:"estado,omitempty" bson:"estado"`
	Referencia     string       `json:"referencia,omitempty" bson:"referencia"`
	Directorio     []Directorio `json:"directorio" bson:"directorio"`
}
