package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nikitaenmi/URLShortener/internal/domain"
)

func HTTPErrorHandler(err error, c echo.Context) {
	switch e := err.(type) {
	case *domain.Error:
		//nolint:errcheck
		c.JSON(e.Code, map[string]any{"error": e.Message, "details": e.Details})
	case *echo.HTTPError:
		//nolint:errcheck
		c.JSON(e.Code, e.Message)
	default:
		//nolint:errcheck
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}
}
