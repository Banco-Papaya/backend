package models

import (
	"time"

	"gorm.io/gorm"
)

type Sucursales struct {
	ID        *uint64        `json:"id" gorm:"primary_key;auto_increment"`
	Nombre    string         `json:"nombre" gorm:"type:varchar(255);not null" validate:"required"`
	Direccion string         `json:"direccion" gorm:"type:varchar(255);not null" validate:"required"`
	Telefono  string         `json:"telefono" gorm:"type:varchar(15);not null" validate:"required, min=10, max=15,numeric"`
	Municipio Municipio      `gorm:"foreignkey:ID"`
	CreatedAt time.Time      `json:"createAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index;softDelete"`
}
