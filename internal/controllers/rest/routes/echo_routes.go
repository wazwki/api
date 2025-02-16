package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/wazwki/api/internal/controllers/rest"
)

func RegisterRoutes(e *echo.Echo, controllers rest.NameControllersInterface) {
	api := e.Group("/api")
	v1 := api.Group("/v1")

	name := v1.Group("/name")
	name.GET("/healthcheck", controllers.HealthCheck)
}
