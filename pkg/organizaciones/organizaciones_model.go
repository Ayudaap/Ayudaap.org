package organizaciones

import (
	au "Ayudaap.org/pkg/auditoria"
	di "Ayudaap.org/pkg/direccion"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Organizacion Modelo de datos que representa una organizacion
type Organizacion struct {
	ID        primitive.ObjectID `json:"Id" bson:"_id, omitempty"`
	Tipo      TipoOrganizacion   `json:"tipo,omitempty" bson:"tipo"`
	Nombre    string             `json:"nombre,omitempty" bson:"nombre"`
	Domicilio di.Directorio      `json:"direccion,omitempty" bson:"direccion"`
	Auditoria au.Auditoria       `json:"auditoria,omitempty" bson:"auditoria"`
	Banner    string             `json:"banner,omitempty" bson:"banner"`
}

// TipoOrganizacion Tipo de organizacion
type TipoOrganizacion int

const (

	//OrganizacionGubernamental Tipo de organizacion Gubernamental
	OrganizacionGubernamental TipoOrganizacion = iota
	//OrganizacionNoGubernamental Organizacion no gubernamental
	OrganizacionNoGubernamental
	//OrganizacionSocialConFinesDeLucro Organizacion social con fines de lucro
	OrganizacionSocialConFinesDeLucro
	//OrganizacionSocialSinFinesDeLucro Organizacion social sin fines de lucro
	OrganizacionSocialSinFinesDeLucro
	//OrganizacionPrivada Organizacion privada
	OrganizacionPrivada
)
