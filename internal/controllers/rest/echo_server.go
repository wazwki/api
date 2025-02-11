package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/wazwki/skillsrock/internal/config"
	"github.com/wazwki/skillsrock/internal/controllers/rest/middlewares"
	"github.com/wazwki/skillsrock/pkg/jwtutil"
)

func NewEchoServer(cfg *config.Config, jwt *jwtutil.JWTUtil) *echo.Echo {
	srv := echo.New()
	srv.HideBanner = true
	srv.GET("/swagger/*", echoSwagger.WrapHandler)
	srv.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	srv.Use(
		echo.MiddlewareFunc(middlewares.MetricsMiddleware()),
		echo.MiddlewareFunc(middlewares.LoggerMiddleware()),
	)
	if !cfg.Debug {
		srv.Use(
			echo.MiddlewareFunc(middlewares.JWTMiddleware(cfg, jwt)),
		)
	}

	srv.Server = &http.Server{
		Addr:              fmt.Sprintf("%v:%v", cfg.Host, cfg.Port),
		ReadHeaderTimeout: 1000 * time.Millisecond,
		ReadTimeout:       1000 * time.Millisecond,
	}

	return srv
}
