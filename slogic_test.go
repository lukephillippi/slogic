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

func mockFilter(result bool) Filter {
	return func(ctx context.Context, r slog.Record) bool {
		return result
	}
}
