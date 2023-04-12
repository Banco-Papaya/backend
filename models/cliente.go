package models

import (
	"time"

	"gorm.io/gorm"
)

type Cliente struct {
	ID              *uint64        `json:"id" gorm:"primary_key;auto_increment"`
	NombreCompleto  string         `json:"nombreCompleto" gorm:"type:varchar(255);not null" validate:"required"`
	Correo          string         `json:"Correo" gorm:"type:varchar(255);not null" validate:"required, min=6, max=255, email"`
	Telefono        string         `json:"telefono" gorm:"type:varchar(15);not null" validate:"required, min=10, max=15, numeric"`
	Direccion       string         `json:"direccion" gorm:"type:varchar(255);not null" validate:"required, min=1, max=255"`
	Estado          bool           `json:"estado" gorm:"default:true"`
	Clave           string         `json:"clave" gorm:"type:varchar(255);not null" validate:"required"`
	NumeroDocumento string         `json:"numeroDocumento" gorm:"type:varchar(15);not null" validate:"required, min=10, max=15, numeric"`
	TipoDocumento   TipoDocumento  `json:"tipoDocumento" gorm:"foreignkey:ID"`
	Municipio       Municipio      `json:"municipio" gorm:"foreignkey:ID"`
	Sucursal        Sucursales     `json:"sucursal" gorm:"foreignkey:ID"`
	CreatedAt       time.Time      `json:"createAt" gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `json:"deletedAt" gorm:"index;softDelete"`
}
