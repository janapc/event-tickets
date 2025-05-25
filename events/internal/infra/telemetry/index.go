package telemetry

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

func InitOpenTelemetry(ctx context.Context) func() {
	res, err := resource.New(ctx, resource.WithFromEnv())
	if err != nil {
		log.Fatalf("failed to create resource: %v", err)
	}

	// Tracer
	tp := initTracer(res, ctx)
	otel.SetTracerProvider(tp)
	// Metrics
	mp := initMetrics(res, ctx)
	otel.SetMeterProvider(mp)

	return func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := tp.Shutdown(shutdownCtx); err != nil {
			log.Printf("error shutting down tracer provider: %v", err)
		}
		if err := mp.Shutdown(shutdownCtx); err != nil {
			log.Printf("error shutting down meter provider: %v", err)
		}
	}
}

func initTracer(res *resource.Resource, ctx context.Context) *trace.TracerProvider {
	tracerExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to create trace exporter: %v", err)
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
		log.Fatalf("failed to create trace exporter: %v", err)
	}
	mp := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(metricExporter)),
		metric.WithResource(res),
	)
	return mp
}
