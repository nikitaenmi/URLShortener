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
		HandleError: true,

		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			ctx := c.Request().Context()
			logWithContext := WithContext(ctx, log)

			if v.Error == nil {
				logWithContext.Info("HTTP request",
					"method", v.Method,
					"uri", v.URI,
					"status", v.Status,
					"latency", v.Latency,
					"remote_ip", v.RemoteIP,
				)
			} else {
				logWithContext.Error("HTTP request error",
					"method", v.Method,
					"uri", v.URI,
					"status", v.Status,
					"error", v.Error.Error(),
					"latency", v.Latency,
				)
			}
			return nil
		},
	})
}
