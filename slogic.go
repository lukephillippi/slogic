// Package slogic provides composable filtering logic for [log/slog].
package slogic // import "go.luke.ph/slogic"

import (
	"context"
	"log/slog"
)

var _ slog.Handler = (*Handler)(nil)

// A Filter returns true if the given [slog.Record] should be filtered out,
// and returns false if not.
type Filter func(context.Context, slog.Record) bool

// NewHandler constructs a [*Handler] that wraps the given handler with a filter.
func NewHandler(handler slog.Handler, filter Filter) *Handler {
	return &Handler{
		handler: handler,
		filter:  filter,
	}
}

// A Handler implements the [slog.Handler] interface.
type Handler struct {
	handler slog.Handler
	filter  Filter
}

// Enabled implements the [slog.Handler] Enabled interface method.
// It calls the wrapped handler's Enabled method, unaffected by the filter.
func (h *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

// Handle implements the [slog.Handler] Handle interface method.
// It calls the wrapped handler's Handle method only if the filter returns false.
func (h *Handler) Handle(ctx context.Context, r slog.Record) error {
	if h.filter(ctx, r) {
		return nil
	}
	return h.handler.Handle(ctx, r)
}

// WithAttrs implements the [slog.Handler] WithAttrs interface method.
// It calls the wrapped handler's WithAttrs method, unaffected by the filter.
func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Handler{
		handler: h.handler.WithAttrs(attrs),
		filter:  h.filter,
	}
}

// WithGroup implements the [slog.Handler] WithGroup interface method.
// It calls the wrapped handler's WithGroup method, unaffected by the filter.
func (h *Handler) WithGroup(name string) slog.Handler {
	return &Handler{
		handler: h.handler.WithGroup(name),
		filter:  h.filter,
	}
}
