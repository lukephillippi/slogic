package filter

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"go.luke.ph/slogic"
)

func TestIfTimeAfter(t *testing.T) {
	now := time.Now()
	before := now.Add(-time.Second)
	after := now.Add(time.Second)

	tests := []struct {
		name   string
		record time.Time
		filter time.Time
		want   bool
	}{
		{
			name:   "equal",
			record: now,
			filter: now,
			want:   false,
		},
		{
			name:   "before",
			record: now,
			filter: after,
			want:   false,
		},
		{
			name:   "after",
			record: now,
			filter: before,
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := testTime(IfTimeAfter(tt.filter), tt.record)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestIfTimeBefore(t *testing.T) {
	now := time.Now()
	before := now.Add(-time.Second)
	after := now.Add(time.Second)

	tests := []struct {
		name   string
		record time.Time
		filter time.Time
		want   bool
	}{
		{
			name:   "equal",
			record: now,
			filter: now,
			want:   false,
		},
		{
			name:   "before",
			record: now,
			filter: after,
			want:   true,
		},
		{
			name:   "after",
			record: now,
			filter: before,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := testTime(IfTimeBefore(tt.filter), tt.record)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestIfTimeBetween(t *testing.T) {
	now := time.Now()
	before := now.Add(-time.Second)
	after := now.Add(time.Second)

	tests := []struct {
		name   string
		record time.Time
		start  time.Time
		end    time.Time
		want   bool
	}{
		{
			name:   "in range",
			record: now,
			start:  before,
			end:    after,
			want:   true,
		},
		{
			name:   "equal to start",
			record: now,
			start:  now,
			end:    after,
			want:   true,
		},
		{
			name:   "equal to end",
			record: now,
			start:  before,
			end:    now,
			want:   true,
		},
		{
			name:   "before range",
			record: before,
			start:  now,
			end:    after,
			want:   false,
		},
		{
			name:   "after range",
			record: after,
			start:  before,
			end:    now,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := testTime(IfTimeBetween(tt.start, tt.end), tt.record)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func testTime(filter slogic.Filter, time time.Time) bool {
	return filter(
		context.Background(),
		slog.NewRecord(time, slog.LevelInfo, "", 0),
	)
}
