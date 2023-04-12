package routes

import (
	"myapp/controllers"

	"github.com/labstack/echo/v4"
)

func SetDepartametos(router *echo.Echo) {
	subRoute := router.Group("/departamento/api")
	subRoute.GET("/all", func(c echo.Context) error {
		return controllers.GetAllDepartamentos(c)
	})
}
