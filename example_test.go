package slogic_test

import (
	"log/slog"
	"os"
	"time"

	"go.luke.ph/slogic"
	"go.luke.ph/slogic/filter"
)

func Example() {
	// 1. Keep all ERROR logs
	// 2. Keep latency WARN logs
	// 3. Filter all other ≤ WARN logs
	handler := slogic.NewHandler(
		slog.NewTextHandler(os.Stdout, opts),
		slogic.And(
			slogic.Not(
				slogic.Or(
					filter.IfLevelEquals(slog.LevelError),
					slogic.And(
						filter.IfLevelEquals(slog.LevelWarn),
						filter.IfAttrExists("latency_ms"),
					),
				),
			),
			filter.IfLevelAtMost(slog.LevelWarn),
		),
	)

	logger := slog.New(handler)

	logger.Debug("Received request", "method", "GET", "path", "/api/users", "ip", "192.168.1.1") // Filtered
	logger.Info("Authenticated user", "user_id", "user_123", "roles", "admin,reader")            // Filtered
	logger.Warn("Executed slow database query", "query", "getUserProfile", "latency_ms", 250)
	logger.Error("Failed to process payment", "order_id", "ORD-9876", "error", "gateway_timeout")

	// Output:
	// time=1970-01-01T00:00:00.000Z level=WARN msg="Executed slow database query" query=getUserProfile latency_ms=250
	// time=1970-01-01T00:00:00.000Z level=ERROR msg="Failed to process payment" order_id=ORD-9876 error=gateway_timeout
}

func ExampleAnd() {
	handler := slogic.NewHandler(
		slog.NewTextHandler(os.Stdout, opts),
		slogic.And(
			filter.IfMessageContains("database"),
			filter.IfAttrExists("latency_ms"),
		),
	)

	logger := slog.New(handler)

	logger.Debug("Received request", "method", "GET", "path", "/api/users", "ip", "192.168.1.1")
	logger.Info("Authenticated user", "user_id", "user_123", "roles", "admin,reader")
	logger.Warn("Executed slow database query", "query", "getUserProfile", "latency_ms", 250) // Filtered
	logger.Error("Failed to process payment", "order_id", "ORD-9876", "error", "gateway_timeout")

	// Output:
	// time=1970-01-01T00:00:00.000Z level=DEBUG msg="Received request" method=GET path=/api/users ip=192.168.1.1
	// time=1970-01-01T00:00:00.000Z level=INFO msg="Authenticated user" user_id=user_123 roles=admin,reader
	// time=1970-01-01T00:00:00.000Z level=ERROR msg="Failed to process payment" order_id=ORD-9876 error=gateway_timeout
}

func ExampleOr() {
	handler := slogic.NewHandler(
		slog.NewTextHandler(os.Stdout, opts),
		slogic.Or(
			filter.IfLevelEquals(slog.LevelDebug),
			filter.IfAttrContains("roles", "admin"),
		),
	)

	logger := slog.New(handler)

	logger.Debug("Received request", "method", "GET", "path", "/api/users", "ip", "192.168.1.1") // Filtered
	logger.Info("Authenticated user", "user_id", "user_123", "roles", "admin,reader")            // Filtered
	logger.Warn("Executed slow database query", "query", "getUserProfile", "latency_ms", 250)
	logger.Error("Failed to process payment", "order_id", "ORD-9876", "error", "gateway_timeout")

	// Output:
	// time=1970-01-01T00:00:00.000Z level=WARN msg="Executed slow database query" query=getUserProfile latency_ms=250
	// time=1970-01-01T00:00:00.000Z level=ERROR msg="Failed to process payment" order_id=ORD-9876 error=gateway_timeout
}

func ExampleNot() {
	handler := slogic.NewHandler(
		slog.NewTextHandler(os.Stdout, opts),
		slogic.Not(
			filter.IfMessageContains("payment"),
		),
	)

	logger := slog.New(handler)

	logger.Debug("Received request", "method", "GET", "path", "/api/users", "ip", "192.168.1.1") // Filtered
	logger.Info("Authenticated user", "user_id", "user_123", "roles", "admin,reader")            // Filtered
	logger.Warn("Executed slow database query", "query", "getUserProfile", "latency_ms", 250)    // Filtered
	logger.Error("Failed to process payment", "order_id", "ORD-9876", "error", "gateway_timeout")

	// Output:
	// time=1970-01-01T00:00:00.000Z level=ERROR msg="Failed to process payment" order_id=ORD-9876 error=gateway_timeout
}

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
