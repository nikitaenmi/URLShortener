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

			// 404
			if errors.Is(err, domain.ErrURLNotFound) {
				log.Info("URL not found" /* ... */)
				return c.JSON(http.StatusNotFound, map[string]string{"error": "URL not found"})
			}

			// 400 — клиентские ошибки
			if errors.Is(err, domain.ErrInvalidRequest) ||
				errors.Is(err, domain.ErrInvalidID) ||
				errors.Is(err, domain.ErrInvalidQueryParams) {
				log.Warn("Client error",
					"error", err.Error(),
					"method", c.Request().Method,
					"path", c.Request().URL.Path,
					"request_id", requestID,
				)
				// Возвращаем сообщение из самой ошибки
				return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			}

			// 500 — всё остальное
			log.Error("Internal server error" /* ... */)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
		}
	}
}
