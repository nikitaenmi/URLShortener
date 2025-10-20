package middleware

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nikitaenmi/URLShortener/internal/domain"
	"github.com/nikitaenmi/URLShortener/internal/lib/logger"
)

func ErrorHandler(log logger.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err == nil {
				return nil
			}

			requestID := GetRequestIDFromContext(c.Request().Context())

			if errors.Is(err, domain.ErrURLNotFound) {
				return c.JSON(http.StatusNotFound, map[string]string{
					"error": "URL not found",
				})
			}

			log.Error(
				"Internal server error",
				"error", err,
				"method", c.Request().Method,
				"path", c.Request().URL.Path,
				"remote_ip", c.Request().RemoteAddr,
				"request_id", requestID,
			)

			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Internal server error",
			})
		}
	}
}
