package logger

import (
	"context"

	"github.com/nikitaenmi/URLShortener/internal/http-server/middleware"
)

type ContextLogger struct {
	logger Logger
	ctx    context.Context
}

func WithContext(ctx context.Context, logger Logger) *ContextLogger {
	return &ContextLogger{
		logger: logger,
		ctx:    ctx,
	}
}

func (cl *ContextLogger) Info(msg string, args ...any) {
	cl.logger.Info(msg, cl.appendRequestID(args)...)
}

func (cl *ContextLogger) Error(msg string, args ...any) {
	cl.logger.Error(msg, cl.appendRequestID(args)...)
}

func (cl *ContextLogger) Warn(msg string, args ...any) {
	cl.logger.Warn(msg, cl.appendRequestID(args)...)
}

func (cl *ContextLogger) appendRequestID(args []any) []any {
	if reqID := middleware.GetRequestIDFromContext(cl.ctx); reqID != "" {
		return append(args, "request_id", reqID)
	}
	return args
}
