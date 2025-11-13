package middleware

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const RequestIDKey = "request_id"

func RequestIDMiddleware() echo.MiddlewareFunc {
	config := middleware.RequestIDConfig{
		RequestIDHandler: func(c echo.Context, requestID string) {
			ctx := context.WithValue(c.Request().Context(), RequestIDKey, requestID)
			c.SetRequest(c.Request().WithContext(ctx))
		},
	}
	return middleware.RequestIDWithConfig(config)
}

func GetRequestIDFromContext(ctx context.Context) string {
	if requestID, ok := ctx.Value(RequestIDKey).(string); ok {
		return requestID
	}
	return ""
}
