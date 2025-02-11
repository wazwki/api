package rest

import (
	"github.com/labstack/echo/v4"
)

type NameControllersInterface interface {
	HealthCheck(c echo.Context) error
}
