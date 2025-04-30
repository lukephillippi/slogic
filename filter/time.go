package filter

import (
	"context"
	"log/slog"
	"time"

	"go.luke.ph/slogic"
)

// IfTimeAfter returns a [slogic.Filter] that returns true if
// the record's Time is after the given time.
func IfTimeAfter(time time.Time) slogic.Filter {
	return func(_ context.Context, r slog.Record) bool {
		return r.Time.After(time)
	}
}

// IfTimeBefore returns a [slogic.Filter] that returns true if
// the record's Time is before the given time.
func IfTimeBefore(time time.Time) slogic.Filter {
	return func(_ context.Context, r slog.Record) bool {
		return r.Time.Before(time)
	}
}

// IfTimeBetween returns a [slogic.Filter] that returns true if
// the record's Time is between the given start and end times.
func IfTimeBetween(start, end time.Time) slogic.Filter {
	return func(_ context.Context, r slog.Record) bool {
		return !r.Time.Before(start) && !r.Time.After(end)
	}
}
