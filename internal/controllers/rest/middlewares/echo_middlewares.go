package middlewares

import (
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/wazwki/skillsrock/internal/config"
	"github.com/wazwki/skillsrock/pkg/jwtutil"
	"github.com/wazwki/skillsrock/pkg/logger"
	"github.com/wazwki/skillsrock/pkg/metrics"
	"go.uber.org/zap"
)

func JWTMiddleware(cfg *config.Config, jwt *jwtutil.JWTUtil) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Path() == "/api/v1/name/healthcheck" {
				return next(c)
			}

			if c.Path() != "/api/v1/name/healthcheck" {
				token, err := jwt.GenerateAccessToken(c.Request().Context())
				if err != nil {
					return err
				}

				c.SetCookie(&http.Cookie{
					Name:     "Authorization",
					Value:    token,
					Expires:  time.Now().Add(time.Second * time.Duration(cfg.AccessTokenTTL)),
					HttpOnly: true,
					SameSite: http.SameSiteLaxMode,
				})

				return next(c)
			}

			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.ErrUnauthorized
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			_, err := jwt.ValidateToken(c.Request().Context(), tokenStr)
			if err != nil {
				return echo.ErrUnauthorized
			}

			return next(c)
		}
	}
}

func MetricsMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			if err != nil {
				return err
			}

			metrics.ObserveRequestDuration.WithLabelValues(c.Request().Method, c.Path()).Observe(time.Since(start).Seconds())

			return nil
		}
	}
}

func LoggerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			logger.Info(c.Request().Method, zap.String("Path", c.Path()), zap.String("module", "skillsrock"))

			err := next(c)
			if err != nil {
				if err != echo.ErrNotFound && err != echo.ErrUnauthorized {
					logger.Error(err.Error(), zap.String("module", "skillsrock"))
				}
				return err
			}
			return nil
		}
	}
}
