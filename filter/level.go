package filter

import (
	"context"
	"log/slog"

	"go.luke.ph/slogic"
)

// IfLevelEquals returns a [slogic.Filter] that returns true if
// the record's Level is equivalent to the given level.
func IfLevelEquals(level slog.Level) slogic.Filter {
	return func(_ context.Context, r slog.Record) bool {
		return r.Level == level
	}
}

// IfLevelAtLeast returns a [slogic.Filter] that returns true if
// the record's Level is at least the given level.
func IfLevelAtLeast(level slog.Level) slogic.Filter {
	return func(_ context.Context, r slog.Record) bool {
		return r.Level >= level
	}
}

// IfLevelAtMost returns a [slogic.Filter] that returns true if
// the record's Level is at most the given level.
func IfLevelAtMost(level slog.Level) slogic.Filter {
	return func(_ context.Context, r slog.Record) bool {
		return r.Level <= level
	}
}
