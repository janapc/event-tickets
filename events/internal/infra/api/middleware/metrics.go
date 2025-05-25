package middleware

import (
	"net/http"
	"strconv"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var (
	meter             = otel.Meter("http-metrics")
	requestLatency, _ = meter.Float64Histogram("http_server_duration", metric.WithUnit("s"), metric.WithDescription("http request latency"))
	requestCount, _   = meter.Int64Counter("http_server_request_count", metric.WithDescription("http request count"))
)

func Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ctx := r.Context()

		rr := &statusObserver{ResponseWriter: w, statusCode: 0}
		next.ServeHTTP(rr, r)

		duration := time.Since(start).Seconds()
		status := rr.statusCode
		if status == 0 {
			status = http.StatusOK
		}

		attrs := []attribute.KeyValue{
			attribute.String("http.method", r.Method),
			attribute.String("http.path", r.URL.Path),
			attribute.String("http.status_code", strconv.Itoa(status)),
		}

		requestCount.Add(ctx, 1, metric.WithAttributes(attrs...))
		requestLatency.Record(ctx, duration, metric.WithAttributes(attrs...))
	})
}

type statusObserver struct {
	http.ResponseWriter
	statusCode int
}

func (s *statusObserver) WriteHeader(code int) {
	s.statusCode = code
	s.ResponseWriter.WriteHeader(code)
}
