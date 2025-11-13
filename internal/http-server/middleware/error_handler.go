package middleware

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nikitaenmi/URLShortener/internal/domain"
)

func HTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	code := http.StatusInternalServerError
	message := "Internal Server Error"

	switch {
	case errors.Is(err, domain.ErrURLNotFound):
		code = http.StatusNotFound
		message = "url not found"

	case errors.Is(err, domain.ErrInvalidID):
		code = http.StatusBadRequest
		message = "invalid ID format"

	case errors.Is(err, domain.ErrInvalidRequest):
		code = http.StatusBadRequest
		message = "invalid request"

	case errors.Is(err, domain.ErrInvalidQueryParams):
		code = http.StatusBadRequest
		message = "invalid query parameters"

	default:
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			if msg, ok := he.Message.(string); ok {
				message = msg
			}
		}
	}

	_ = c.JSON(code, map[string]string{"error": message})
}
