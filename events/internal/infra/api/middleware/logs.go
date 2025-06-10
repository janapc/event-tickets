package middleware

import (
	"net/http"

	"github.com/janapc/event-tickets/events/internal/infra/logger"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func RequestTracerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		span := trace.SpanFromContext(ctx)

		if span.IsRecording() {
			span.SetAttributes(
				attribute.String("http.request.method", r.Method),
				attribute.String("http.request.path", r.URL.Path),
			)
			logger.Logger.WithContext(ctx).WithFields(logrus.Fields{
				"http.method": r.Method,
				"http.path":   r.URL.Path,
			}).Info("Request method and path added to span and logs.")
		}

		next.ServeHTTP(w, r)
	})
}
