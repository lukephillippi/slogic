package filter_test

import (
	"log/slog"
	"time"
)

var opts = &slog.HandlerOptions{
	Level: slog.LevelDebug,
	// Replaces the log time with a fixed value for testable examples...
	ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey {
			a.Value = slog.TimeValue(time.Unix(0, 0).UTC())
		}
		return a
	},
}
