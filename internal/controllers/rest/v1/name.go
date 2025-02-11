package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wazwki/skillsrock/internal/controllers/rest"
	"github.com/wazwki/skillsrock/internal/service"
)

type NameControllers struct {
	service service.NameServiceInterface
}

func NewNameControllers(s service.NameServiceInterface) rest.NameControllersInterface {
	return &NameControllers{service: s}
}

// @Summary Healthcheck
// @Description Healthcheck
// @Accept json
// @Produce json
// @Success 201 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/name/healthcheck [get]
func (controllers *NameControllers) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
