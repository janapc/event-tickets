package middleware

import (
	"net/http"
	"time"

	"github.com/janapc/event-tickets/events/internal/infra/logger"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(rec, r)
		duration := time.Since(start).Milliseconds()
		logger.Logger.WithContext(ctx).Infof("[HTTP] Request: %s %s - Status: %d - Response Time: %vms\n", r.Method, r.URL.Path, rec.status, duration)
	})
}
