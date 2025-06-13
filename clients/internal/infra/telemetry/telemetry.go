package telemetry

import (
	"context"
	"fmt"
	"log/slog"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

var TracerProvider *sdktrace.TracerProvider
var Tracer trace.Tracer

func Init(ctx context.Context) error {
	res, err := resource.New(ctx, resource.WithFromEnv())
	if err != nil {
		return fmt.Errorf("failed to create resource: %w", err)
	}

	// Tracer
	err = initTracer(res, ctx)
	if err != nil {
		return err
	}
	otel.SetTracerProvider(TracerProvider)
	slog.Info("OpenTelemetry initialized successfully.")
	return nil
}

func initTracer(res *resource.Resource, ctx context.Context) error {
	tracerExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to create tracer exporter: %w", err)
	}
	TracerProvider = sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(tracerExporter),
		sdktrace.WithResource(res),
	)
	Tracer = otel.Tracer("clients-service")
	return nil
}

func Shutdown(ctx context.Context) error {
	if TracerProvider != nil {
		if err := TracerProvider.Shutdown(ctx); err != nil {
			slog.Error("Failed to shutdown OpenTelemetry TracerProvider")
			return fmt.Errorf("failed to shutdown TracerProvider: %w", err)
		}
		slog.Info("OpenTelemetry TracerProvider shut down.")
	}
	return nil
}
