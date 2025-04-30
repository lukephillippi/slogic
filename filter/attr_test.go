package filter

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"go.luke.ph/slogic"
)

func TestIfAttrEquals(t *testing.T) {
	tests := []struct {
		name  string
		attrs []slog.Attr
		want  bool
	}{
		{
			name:  "false",
			attrs: []slog.Attr{slog.String("FOO", "BAZ")},
			want:  false,
		},
		{
			name:  "true",
			attrs: []slog.Attr{slog.String("FOO", "BAR")},
			want:  true,
		},
		{
			name:  "missing",
			attrs: []slog.Attr{slog.String("BAR", "BAZ")},
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := testAttr(IfAttrEquals("FOO", "BAR"), tt.attrs)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestIfAttrContains(t *testing.T) {
	tests := []struct {
		name   string
		substr string
		attrs  []slog.Attr
		want   bool
	}{
		{
			name:   "false",
			substr: "fail",
			attrs:  []slog.Attr{slog.String("FOO", "successful")},
			want:   false,
		},
		{
			name:   "true",
			substr: "fail",
			attrs:  []slog.Attr{slog.String("FOO", "failed")},
			want:   true,
		},
		{
			name:   "missing",
			substr: "fail",
			attrs:  []slog.Attr{slog.String("BAR", "failed")},
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := testAttr(IfAttrContains("FOO", tt.substr), tt.attrs)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestIfAttrMatches(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		attrs   []slog.Attr
		want    bool
	}{
		{
			name:    "false",
			pattern: `^user-\d+$`,
			attrs:   []slog.Attr{slog.String("FOO", "admin-123")},
			want:    false,
		},
		{
			name:    "true",
			pattern: `^user-\d+$`,
			attrs:   []slog.Attr{slog.String("FOO", "user-123")},
			want:    true,
		},
		{
			name:    "missing",
			pattern: `^user-\d+$`,
			attrs:   []slog.Attr{slog.String("BAR", "user-123")},
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := testAttr(IfAttrMatches("FOO", tt.pattern), tt.attrs)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestIfAttrExists(t *testing.T) {
	tests := []struct {
		name  string
		attrs []slog.Attr
		want  bool
	}{
		{
			name:  "false",
			attrs: []slog.Attr{slog.String("BAR", "BAZ")},
			want:  false,
		},
		{
			name:  "true",
			attrs: []slog.Attr{slog.String("FOO", "BAR")},
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := testAttr(IfAttrExists("FOO"), tt.attrs)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func testAttr(filter slogic.Filter, attrs []slog.Attr) bool {
	r := slog.NewRecord(time.Now(), slog.LevelInfo, "", 0)
	for _, attr := range attrs {
		r.AddAttrs(attr)
	}
	return filter(context.Background(), r)
}
