package filter_test

import (
	"log/slog"
	"os"

	"go.luke.ph/slogic"
	"go.luke.ph/slogic/filter"
)

func ExampleIfMessageEquals() {
	handler := slogic.NewHandler(
		slog.NewTextHandler(os.Stdout, opts),
		filter.IfMessageEquals("Authenticated user"),
	)

	logger := slog.New(handler)

	logger.Debug("Received request", "method", "GET", "path", "/api/users", "ip", "192.168.1.1")
	logger.Info("Authenticated user", "user_id", "user_123", "roles", "admin,reader") // Filtered
	logger.Warn("Executed slow database query", "query", "getUserProfile", "latency_ms", 250)
	logger.Error("Failed to process payment", "order_id", "ORD-9876", "error", "gateway_timeout")

	// Output:
	// time=1970-01-01T00:00:00.000Z level=DEBUG msg="Received request" method=GET path=/api/users ip=192.168.1.1
	// time=1970-01-01T00:00:00.000Z level=WARN msg="Executed slow database query" query=getUserProfile latency_ms=250
	// time=1970-01-01T00:00:00.000Z level=ERROR msg="Failed to process payment" order_id=ORD-9876 error=gateway_timeout
}

func ExampleIfMessageContains() {
	handler := slogic.NewHandler(
		slog.NewTextHandler(os.Stdout, opts),
		filter.IfMessageContains("database"),
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

func ExampleIfMessageMatches() {
	handler := slogic.NewHandler(
		slog.NewTextHandler(os.Stdout, opts),
		filter.IfMessageMatches("^Failed.*payment$"),
	)

	logger := slog.New(handler)

	logger.Debug("Received request", "method", "GET", "path", "/api/users", "ip", "192.168.1.1")
	logger.Info("Authenticated user", "user_id", "user_123", "roles", "admin,reader")
	logger.Warn("Executed slow database query", "query", "getUserProfile", "latency_ms", 250)
	logger.Error("Failed to process payment", "order_id", "ORD-9876", "error", "gateway_timeout") // Filtered

	// Output:
	// time=1970-01-01T00:00:00.000Z level=DEBUG msg="Received request" method=GET path=/api/users ip=192.168.1.1
	// time=1970-01-01T00:00:00.000Z level=INFO msg="Authenticated user" user_id=user_123 roles=admin,reader
	// time=1970-01-01T00:00:00.000Z level=WARN msg="Executed slow database query" query=getUserProfile latency_ms=250
}
