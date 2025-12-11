package logger

import (
	"context"
	"fmt"

	"github.com/nikitaenmi/URLShortener/internal/http-server/middleware"
)

type LoggerWithContex struct{}

func NewLoggerWithContex() *LoggerWithContex {
	return &LoggerWithContex{}
}

func (l *LoggerWithContex) Info(ctx context.Context, msg string, args ...any) {
	args = appendRequestIDAndArgs(ctx, args)
	fmt.Printf("[INFO] %s %v\n", msg, args)
}

func (l *LoggerWithContex) Error(ctx context.Context, msg string, args ...any) {
	args = appendRequestIDAndArgs(ctx, args)
	fmt.Printf("[ERROR] %s %v\n", msg, args)
}

func (l *LoggerWithContex) Warn(ctx context.Context, msg string, args ...any) {
	args = appendRequestIDAndArgs(ctx, args)
	fmt.Printf("[WARN] %s %v\n", msg, args)
}

func appendRequestIDAndArgs(ctx context.Context, args []any) []any {
	if reqID := middleware.GetRequestIDFromContext(ctx); reqID != "" {
		return append([]any{"request_id", reqID}, args...)
	}
	return args
}
