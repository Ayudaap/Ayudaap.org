package entities

// Tipo de organizacion
type TipoOrganizacion int

const (
	OrganizacionGubernamental TipoOrganizacion = iota
	OrganizacionNoGubernamental
	OrganizacionSocialConFinesDeLucro
	OrganizacionSocialSinFinesDeLucro
	OrganizacionPrivada
)
