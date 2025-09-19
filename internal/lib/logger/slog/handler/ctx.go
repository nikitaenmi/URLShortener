package handler

import (
	"context"
	"log/slog"
)

const RequestIDLogKey = "request_id"

type CtxHandler struct {
	Next slog.Handler
}

func NewCtxHandler(handler slog.Handler) slog.Handler {
	return CtxHandler{Next: handler}
}

// Handle - обработка записи лога с добавлением данных из контекста
func (h CtxHandler) Handle(ctx context.Context, r slog.Record) error {
	if reqID, ok := ctx.Value(RequestIDLogKey).(string); ok {
		r.AddAttrs(slog.String(RequestIDLogKey, reqID))
	}

	return h.Next.Handle(ctx, r)
}

func (h CtxHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h.Next.WithAttrs(attrs)
}

func (h CtxHandler) WithGroup(name string) slog.Handler {
	return h.Next.WithGroup(name)
}

func (h CtxHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.Next.Enabled(ctx, level)
}
