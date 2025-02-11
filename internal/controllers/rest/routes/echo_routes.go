package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/wazwki/skillsrock/internal/controllers/rest"
)

func RegisterRoutes(e *echo.Echo, controllers rest.NameControllersInterface) {
	api := e.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/name/healthcheck", controllers.HealthCheck)

}
