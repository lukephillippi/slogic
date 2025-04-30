package filter

import (
	"context"
	"log/slog"
	"regexp"
	"strings"

	"go.luke.ph/slogic"
)

// IfMessageEquals returns a [slogic.Filter] that returns true if
// the record's Message is equivalent to the given message.
func IfMessageEquals(message string) slogic.Filter {
	return func(_ context.Context, r slog.Record) bool {
		return r.Message == message
	}
}

// IfMessageContains returns a [slogic.Filter] that returns true if
// the record's Message contains the given substring.
func IfMessageContains(substring string) slogic.Filter {
	return func(_ context.Context, r slog.Record) bool {
		return strings.Contains(r.Message, substring)
	}
}

// IfMessageMatches returns a [slogic.Filter] that returns true if
// the record's Message matches the given regular expression.
func IfMessageMatches(pattern string) slogic.Filter {
	re := regexp.MustCompile(pattern)
	return func(_ context.Context, r slog.Record) bool {
		return re.MatchString(r.Message)
	}
}
