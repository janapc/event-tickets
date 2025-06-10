package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

var Logger *logrus.Logger

type TraceContextHook struct{}

func (h *TraceContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *TraceContextHook) Fire(entry *logrus.Entry) error {
	span := trace.SpanFromContext(entry.Context)
	spanCtx := span.SpanContext()

	if spanCtx.IsValid() {
		entry.Data["trace.id"] = spanCtx.TraceID().String()
		entry.Data["span.id"] = spanCtx.SpanID().String()
		entry.Data["trace.flags"] = fmt.Sprintf("%02x", spanCtx.TraceFlags())
	}
	return nil
}

func Init() {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.JSONFormatter{})
	Logger.SetOutput(os.Stdout)
	Logger.SetLevel(logrus.InfoLevel)
	Logger.AddHook(&TraceContextHook{})
}
