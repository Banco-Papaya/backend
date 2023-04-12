package common

import (
	"log"
	"myapp/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@/bancopapaya?charset=utf8")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()

	log.Println("Migrando.....")

	db.AutoMigrate(&models.Municipio{})
	db.AutoMigrate(&models.Departamento{})
	db.AutoMigrate(&models.TipoDocumento{})
	db.AutoMigrate(&models.Cliente{})
	db.AutoMigrate(&models.Rol{})
	db.AutoMigrate(&models.Empleado{})
	db.AutoMigrate(&models.Sucursales{})

}
