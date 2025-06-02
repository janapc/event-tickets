package telemetry

import (
	"context"
	"log/slog"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"

	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"

	"go.opentelemetry.io/otel/log/global"

	sdklog "go.opentelemetry.io/otel/sdk/log"
)

func InitOpenTelemetry(ctx context.Context) func() {
	res, err := resource.New(ctx, resource.WithFromEnv())
	if err != nil {
		slog.ErrorContext(ctx, "failed to create resource", "error", err.Error())
	}

	// Tracer
	tp := initTracer(res, ctx)
	otel.SetTracerProvider(tp)
	// Metrics
	mp := initMetrics(res, ctx)
	otel.SetMeterProvider(mp)
	// Logs
	lg := initLogs(res, ctx)
	global.SetLoggerProvider(lg)

	return func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := tp.Shutdown(shutdownCtx); err != nil {
			slog.ErrorContext(ctx, "error shutting down tracer provider", "error", err.Error())
		}
		if err := mp.Shutdown(shutdownCtx); err != nil {
			slog.ErrorContext(ctx, "error shutting down meter provider", "error", err.Error())
		}
	}
}

func initTracer(res *resource.Resource, ctx context.Context) *trace.TracerProvider {
	tracerExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure())
	if err != nil {
		slog.ErrorContext(ctx, "failed to create tracer exporter", "error", err.Error())
	}
	tp := trace.NewTracerProvider(
		trace.WithBatcher(tracerExporter),
		trace.WithResource(res),
	)
	return tp
}

func initMetrics(res *resource.Resource, ctx context.Context) *metric.MeterProvider {
	metricExporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithInsecure())
	if err != nil {
		slog.ErrorContext(ctx, "failed to create metric exporter", "error", err.Error())
	}
	mp := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(metricExporter)),
		metric.WithResource(res),
	)
	return mp
}

func initLogs(res *resource.Resource, ctx context.Context) *sdklog.LoggerProvider {
	exporter, err := otlploggrpc.New(ctx)

	if err != nil {
		slog.ErrorContext(ctx, "failed to create logs exporter", "error", err.Error())
	}
	procesor := sdklog.NewBatchProcessor(exporter)
	provider := sdklog.NewLoggerProvider(
		sdklog.WithProcessor(procesor),
		sdklog.WithResource(res),
	)

	return provider
}
