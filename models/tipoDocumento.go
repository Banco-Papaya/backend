package models

type TipoDocumento struct {
	ID          *uint64 `json:"id" gorm:"primary_key;auto_increment"`
	Nombre      string  `json:"nombre" gorm:"type:varchar(255);not null" validate:"required"`
	Descripcion string  `json:"descripcion" gorm:"type:varchar(255);not null" validate:"required"`
}
