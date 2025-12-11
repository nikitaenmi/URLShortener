package logger

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func EchoRequestLogger(log Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod:   true,
		LogURI:      true,
		LogStatus:   true,
		LogLatency:  true,
		LogRemoteIP: true,
		LogError:    true,
		HandleError: false,

		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			ctx := c.Request().Context()

			args := []any{
				"method", v.Method,
				"uri", v.URI,
				"status", v.Status,
				"latency", v.Latency,
				"remote_ip", v.RemoteIP,
			}

			if v.Error == nil {
				log.Info(ctx, "HTTP request", args...)
			} else {
				errorArgs := append(args, "error", v.Error.Error())
				log.Error(ctx, "HTTP request error", errorArgs...)
			}
			return nil
		},
	})
}
