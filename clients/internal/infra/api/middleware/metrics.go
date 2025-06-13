package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/janapc/event-tickets/clients/internal/infra/telemetry"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

func OtelMetricMiddleware() fiber.Handler {
	requestCounter, _ := telemetry.Meter.Int64Counter(
		"http_server_request_count",
		metric.WithDescription("Count of HTTP requests"),
	)
	latencyHistogram, _ := telemetry.Meter.Float64Histogram(
		"http_server_latency",
		metric.WithUnit("ms"),
		metric.WithDescription("Latency of HTTP requests in milliseconds"),
	)

	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start).Seconds() * 1000 // Convert to milliseconds

		attrs := []attribute.KeyValue{
			attribute.String("http.method", c.Method()),
			attribute.String("http.path", c.Route().Path),
			attribute.Int("http.status_code", c.Response().StatusCode()),
		}

		ctx := c.UserContext()

		requestCounter.Add(ctx, 1, metric.WithAttributes(attrs...))
		latencyHistogram.Record(ctx, duration, metric.WithAttributes(attrs...))

		return err
	}
}
