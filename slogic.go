// Package slogic provides composable filtering logic for [log/slog].
package slogic // import "go.luke.ph/slogic"

import (
	"context"
	"log/slog"
)

var _ slog.Handler = (*Handler)(nil)

// NewHandler constructs a [*Handler] that wraps the given handler.
func NewHandler(handler slog.Handler) *Handler {
	return &Handler{
		handler: handler,
	}
}

// A Handler implements the [slog.Handler] interface.
type Handler struct {
	handler slog.Handler
}

// Enabled implements the [slog.Handler] Enabled interface method.
func (h *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

// Handle implements the [slog.Handler] Handle interface method.
func (h *Handler) Handle(ctx context.Context, r slog.Record) error {
	return h.handler.Handle(ctx, r)
}

// WithAttrs implements the [slog.Handler] WithAttrs interface method.
func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Handler{
		handler: h.handler.WithAttrs(attrs),
	}
}

// WithGroup implements the [slog.Handler] WithGroup interface method.
func (h *Handler) WithGroup(name string) slog.Handler {
	return &Handler{
		handler: h.handler.WithGroup(name),
	}
}
