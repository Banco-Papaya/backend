package main

import (
	"log"
	"myapp/common"
	"myapp/routes"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	common.Migrate()

	e := echo.New()

	cors := middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderAuthorization, echo.HeaderContentType},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
	})
	e.Use(cors)
	routes.SetDepartametos(e)

	server := &http.Server{
		Addr:    ":9000",
		Handler: e,
	}

	log.Println("Servidor ejecutándose sobre el puerto: 9000")
	log.Println(server.ListenAndServe())
}
