package controllers

import (
	"myapp/common"
	"myapp/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllMunicipiosByDepartamento(c echo.Context) error {
	municipios := []models.Municipio{}
	departamentoID := c.Param("departamento")
	db := common.GetConnection()
	defer db.Close()

	db.Where("Departamento LIKE ?", "%"+departamentoID+"%").Find(&municipios)

	if len(municipios) > 0 {
		return c.JSON(http.StatusOK, municipios)
	} else {
		return c.NoContent(http.StatusNotFound)
	}
}

/*
func GetAllMunicipiosByDepartamento(writer http.ResponseWriter, request *http.Request) {

	municipios := models.Municipio{}
	departamentoID := mux.Vars(request)["departamento"]
	db := common.GetConnection()
	defer db.Close()

	db.Where("Departamento LIKE ?", "%"+departamentoID+"%").Find(&municipios)

	if municipios.ID != nil {
		json, _ := json.Marshal(municipios)
		common.SendResponse(writer, http.StatusOK, json)
	} else {
		common.SendError(writer, http.StatusNotFound)
	}
}
*/
