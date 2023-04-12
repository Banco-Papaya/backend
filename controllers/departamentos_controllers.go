package controllers

import (
	"myapp/common"
	"myapp/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllDepartamentos(c echo.Context) error {
	departamentos := []models.Departamento{}
	db := common.GetConnection()
	defer db.Close()
	db.Find(&departamentos)
	return c.JSON(http.StatusOK, departamentos)
}
