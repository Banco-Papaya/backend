package models

type Departamento struct {
	ID     *uint64 `json:"id" gorm:"primary_key;auto_increment"`
	Nombre string  `json:"nombre" gorm:"type:varchar(255);not null" validate:"required"`
}
