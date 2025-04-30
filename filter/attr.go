package filter

import (
	"context"
	"log/slog"
	"regexp"
	"strings"

	"go.luke.ph/slogic"
)

// IfAttrEquals returns a [slogic.Filter] that returns true if
// the record's [slog.Attr] with the given key is equivalent to the given value.
func IfAttrEquals(key string, value any) slogic.Filter {
	return ifAttr(key, func(attr slog.Attr) bool {
		return attr.Value.Equal(slog.AnyValue(value))
	})
}

// IfAttrContains returns a [slogic.Filter] that returns true if
// the record's [slog.Attr] with the given key contains the given substring.
func IfAttrContains(key, substring string) slogic.Filter {
	return ifAttr(key, func(attr slog.Attr) bool {
		return strings.Contains(attr.Value.String(), substring)
	})
}

// IfAttrMatches returns a [slogic.Filter] that returns true if
// the record's [slog.Attr] with the given key matches the given regular expression.
func IfAttrMatches(key, pattern string) slogic.Filter {
	re := regexp.MustCompile(pattern)
	return ifAttr(key, func(attr slog.Attr) bool {
		return re.MatchString(attr.Value.String())
	})
}

// IfAttrExists returns a [slogic.Filter] that returns true if
// the record's [slog.Attr] with the given key exists.
func IfAttrExists(key string) slogic.Filter {
	return ifAttr(key, func(attr slog.Attr) bool {
		return true
	})
}

func ifAttr(key string, predicate func(attr slog.Attr) bool) slogic.Filter {
	return func(_ context.Context, r slog.Record) bool {
		found := false
		r.Attrs(func(attr slog.Attr) bool {
			if attr.Key == key && predicate(attr) {
				found = true
				return false
			}
			return true
		})
		return found
	}
}
