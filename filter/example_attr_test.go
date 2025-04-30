package filter_test

import (
	"log/slog"
	"os"

	"go.luke.ph/slogic"
	"go.luke.ph/slogic/filter"
)

func ExampleIfAttrEquals() {
	handler := slogic.NewHandler(
		slog.NewTextHandler(os.Stdout, opts),
		filter.IfAttrEquals("path", "/api/users"),
	)

	logger := slog.New(handler)

	logger.Debug("Received request", "method", "GET", "path", "/api/users", "ip", "192.168.1.1") // Filtered
	logger.Info("Authenticated user", "user_id", "user_123", "roles", "admin,reader")
	logger.Warn("Executed slow database query", "query", "getUserProfile", "latency_ms", 250)
	logger.Error("Failed to process payment", "order_id", "ORD-9876", "error", "gateway_timeout")

	// Output:
	// time=1970-01-01T00:00:00.000Z level=INFO msg="Authenticated user" user_id=user_123 roles=admin,reader
	// time=1970-01-01T00:00:00.000Z level=WARN msg="Executed slow database query" query=getUserProfile latency_ms=250
	// time=1970-01-01T00:00:00.000Z level=ERROR msg="Failed to process payment" order_id=ORD-9876 error=gateway_timeout
}

func ExampleIfAttrContains() {
	handler := slogic.NewHandler(
		slog.NewTextHandler(os.Stdout, opts),
		filter.IfAttrContains("roles", "admin"),
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

func ExampleIfAttrMatches() {
	handler := slogic.NewHandler(
		slog.NewTextHandler(os.Stdout, opts),
		filter.IfAttrMatches("ip", "^192\\.168\\..*"),
	)

	logger := slog.New(handler)

	logger.Debug("Received request", "method", "GET", "path", "/api/users", "ip", "192.168.1.1") // Filtered
	logger.Info("Authenticated user", "user_id", "user_123", "roles", "admin,reader")
	logger.Warn("Executed slow database query", "query", "getUserProfile", "latency_ms", 250)
	logger.Error("Failed to process payment", "order_id", "ORD-9876", "error", "gateway_timeout")

	// Output:
	// time=1970-01-01T00:00:00.000Z level=INFO msg="Authenticated user" user_id=user_123 roles=admin,reader
	// time=1970-01-01T00:00:00.000Z level=WARN msg="Executed slow database query" query=getUserProfile latency_ms=250
	// time=1970-01-01T00:00:00.000Z level=ERROR msg="Failed to process payment" order_id=ORD-9876 error=gateway_timeout
}

func ExampleIfAttrExists() {
	handler := slogic.NewHandler(
		slog.NewTextHandler(os.Stdout, opts),
		filter.IfAttrExists("latency_ms"),
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
