package filter

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"go.luke.ph/slogic"
)

func TestIfMessageEquals(t *testing.T) {
	tests := []struct {
		name   string
		record string
		filter string
		want   bool
	}{
		{
			name:   "true",
			record: "FOO",
			filter: "FOO",
			want:   true,
		},
		{
			name:   "false",
			record: "BAR",
			filter: "FOO",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := testMessage(IfMessageEquals(tt.filter), tt.record)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestIfMessageContains(t *testing.T) {
	tests := []struct {
		name   string
		record string
		filter string
		want   bool
	}{
		{
			name:   "true",
			record: "FOO",
			filter: "F",
			want:   true,
		},
		{
			name:   "false",
			record: "BAR",
			filter: "F",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := testMessage(IfMessageContains(tt.filter), tt.record)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestIfMessageMatches(t *testing.T) {
	tests := []struct {
		name   string
		record string
		filter string
		want   bool
	}{
		{
			name:   "true",
			record: "FOO",
			filter: "F.*",
			want:   true,
		},
		{
			name:   "false",
			record: "BAR",
			filter: "F.*",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := testMessage(IfMessageMatches(tt.filter), tt.record)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func testMessage(filter slogic.Filter, message string) bool {
	return filter(
		context.Background(),
		slog.NewRecord(time.Unix(0, 0).UTC(), slog.LevelInfo, message, 0),
	)
}
