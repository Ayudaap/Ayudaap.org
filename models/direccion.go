package models

// Direccion de la oficina
type Direccion struct {
	ID             string `json:"ID"`
	Calle          string `json:"calle,omitempty"`
	NumeroInterior string `json:"numeroInterior,omitempty"`
	NumeroExterior string `json:"numeroExterior,omitempty"`
	Colonia        string `json:"colonia,omitempty"`
	CodigoPostal   string `json:"codigoPostal,omitempty"`
	Estado         string `json:"estado,omitempty"`
	Referencia     string `json:"referencia,omitempty"`
}
