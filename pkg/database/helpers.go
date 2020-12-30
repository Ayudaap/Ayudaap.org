package database

import (
	"time"

	"Ayudaap.org/src/entities"
)

//GetAuditoria Obtiene un registro de auditoria
func GetAuditoria(autor string) entities.Auditoria {
	auditoria := entities.Auditoria{
		CreatedAt:     time.Now().Unix(),
		UpdatedAt:     time.Now().Unix(),
		ModificadoPor: autor,
	}
	return auditoria
}
