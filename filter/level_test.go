package filter

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"go.luke.ph/slogic"
)

func TestIfLevelEquals(t *testing.T) {
	tests := []struct {
		name   string
		record slog.Level
		filter slog.Level
		want   bool
	}{
		{
			name:   "true",
			record: slog.LevelInfo,
			filter: slog.LevelInfo,
			want:   true,
		},
		{
			name:   "false",
			record: slog.LevelError,
			filter: slog.LevelInfo,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := testLevel(IfLevelEquals(tt.filter), tt.record)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestIfLevelAtLeast(t *testing.T) {
	tests := []struct {
		name   string
		record slog.Level
		filter slog.Level
		want   bool
	}{
		{
			name:   "==",
			record: slog.LevelInfo,
			filter: slog.LevelInfo,
			want:   true,
		},
		{
			name:   ">=",
			record: slog.LevelError,
			filter: slog.LevelInfo,
			want:   true,
		},
		{
			name:   "<=",
			record: slog.LevelInfo,
			filter: slog.LevelError,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := testLevel(IfLevelAtLeast(tt.filter), tt.record)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestIfLevelAtMost(t *testing.T) {
	tests := []struct {
		name   string
		record slog.Level
		filter slog.Level
		want   bool
	}{
		{
			name:   "==",
			record: slog.LevelInfo,
			filter: slog.LevelInfo,
			want:   true,
		},
		{
			name:   ">=",
			record: slog.LevelError,
			filter: slog.LevelInfo,
			want:   false,
		},
		{
			name:   "<=",
			record: slog.LevelInfo,
			filter: slog.LevelError,
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := testLevel(IfLevelAtMost(tt.filter), tt.record)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func testLevel(filter slogic.Filter, level slog.Level) bool {
	return filter(
		context.Background(),
		slog.NewRecord(time.Unix(0, 0).UTC(), level, "", 0),
	)
}
