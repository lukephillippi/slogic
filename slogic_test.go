package slogic

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"testing"
	"testing/slogtest"
)

func TestHandler(t *testing.T) {
	var buf bytes.Buffer
	h := NewHandler(
		slog.NewJSONHandler(&buf, nil),
		mockFilter(false),
	)

	results := func() []map[string]any {
		var ms []map[string]any
		for line := range bytes.SplitSeq(buf.Bytes(), []byte{'\n'}) {
			if len(line) == 0 {
				continue
			}
			var m map[string]any
			if err := json.Unmarshal(line, &m); err != nil {
				t.Fatal(err)
			}
			ms = append(ms, m)
		}
		return ms
	}

	err := slogtest.TestHandler(h, results)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAnd(t *testing.T) {
	tests := []struct {
		name    string
		filters []Filter
		want    bool
	}{
		{
			name:    "empty",
			filters: []Filter{},
			want:    true,
		},
		{
			name:    "0: false",
			filters: []Filter{mockFilter(false)},
			want:    false,
		},
		{
			name:    "1: true",
			filters: []Filter{mockFilter(true)},
			want:    true,
		},
		{
			name:    "00: false AND false",
			filters: []Filter{mockFilter(false), mockFilter(false)},
			want:    false,
		},
		{
			name:    "01: false AND true",
			filters: []Filter{mockFilter(false), mockFilter(true)},
			want:    false,
		},
		{
			name:    "10: true AND false",
			filters: []Filter{mockFilter(true), mockFilter(false)},
			want:    false,
		},
		{
			name:    "11: true AND true",
			filters: []Filter{mockFilter(true), mockFilter(true)},
			want:    true,
		},
		{
			name:    "000: false AND false AND false",
			filters: []Filter{mockFilter(false), mockFilter(false), mockFilter(false)},
			want:    false,
		},
		{
			name:    "001: false AND false AND true",
			filters: []Filter{mockFilter(false), mockFilter(false), mockFilter(true)},
			want:    false,
		},
		{
			name:    "010: false AND true AND false",
			filters: []Filter{mockFilter(false), mockFilter(true), mockFilter(false)},
			want:    false,
		},
		{
			name:    "011: false AND true AND true",
			filters: []Filter{mockFilter(false), mockFilter(true), mockFilter(true)},
			want:    false,
		},
		{
			name:    "100: true AND false AND false",
			filters: []Filter{mockFilter(true), mockFilter(false), mockFilter(false)},
			want:    false,
		},
		{
			name:    "101: true AND false AND true",
			filters: []Filter{mockFilter(true), mockFilter(false), mockFilter(true)},
			want:    false,
		},
		{
			name:    "110: true AND true AND false",
			filters: []Filter{mockFilter(true), mockFilter(true), mockFilter(false)},
			want:    false,
		},
		{
			name:    "111: true AND true AND true",
			filters: []Filter{mockFilter(true), mockFilter(true), mockFilter(true)},
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := And(tt.filters...)
			got := filter(context.Background(), slog.Record{})
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestOr(t *testing.T) {
	tests := []struct {
		name    string
		filters []Filter
		want    bool
	}{
		{
			name:    "empty",
			filters: []Filter{},
			want:    false,
		},
		{
			name:    "0: false",
			filters: []Filter{mockFilter(false)},
			want:    false,
		},
		{
			name:    "1: true",
			filters: []Filter{mockFilter(true)},
			want:    true,
		},
		{
			name:    "00: false OR false",
			filters: []Filter{mockFilter(false), mockFilter(false)},
			want:    false,
		},
		{
			name:    "01: false OR true",
			filters: []Filter{mockFilter(false), mockFilter(true)},
			want:    true,
		},
		{
			name:    "10: true OR false",
			filters: []Filter{mockFilter(true), mockFilter(false)},
			want:    true,
		},
		{
			name:    "11: true OR true",
			filters: []Filter{mockFilter(true), mockFilter(true)},
			want:    true,
		},
		{
			name:    "000: false OR false OR false",
			filters: []Filter{mockFilter(false), mockFilter(false), mockFilter(false)},
			want:    false,
		},
		{
			name:    "001: false OR false OR true",
			filters: []Filter{mockFilter(false), mockFilter(false), mockFilter(true)},
			want:    true,
		},
		{
			name:    "010: false OR true OR false",
			filters: []Filter{mockFilter(false), mockFilter(true), mockFilter(false)},
			want:    true,
		},
		{
			name:    "011: false OR true OR true",
			filters: []Filter{mockFilter(false), mockFilter(true), mockFilter(true)},
			want:    true,
		},
		{
			name:    "100: true OR false OR false",
			filters: []Filter{mockFilter(true), mockFilter(false), mockFilter(false)},
			want:    true,
		},
		{
			name:    "101: true OR false OR true",
			filters: []Filter{mockFilter(true), mockFilter(false), mockFilter(true)},
			want:    true,
		},
		{
			name:    "110: true OR true OR false",
			filters: []Filter{mockFilter(true), mockFilter(true), mockFilter(false)},
			want:    true,
		},
		{
			name:    "111: true OR true OR true",
			filters: []Filter{mockFilter(true), mockFilter(true), mockFilter(true)},
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := Or(tt.filters...)
			got := filter(context.Background(), slog.Record{})
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestNot(t *testing.T) {
	tests := []struct {
		name   string
		filter Filter
		want   bool
	}{
		{
			name:   "0: false",
			filter: mockFilter(false),
			want:   true,
		},
		{
			name:   "1: true",
			filter: mockFilter(true),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := Not(tt.filter)
			got := filter(context.Background(), slog.Record{})
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func mockFilter(result bool) Filter {
	return func(ctx context.Context, r slog.Record) bool {
		return result
	}
}
