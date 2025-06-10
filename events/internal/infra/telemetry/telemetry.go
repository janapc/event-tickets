package telemetry

import (
	"context"
	"fmt"

	"github.com/janapc/event-tickets/events/internal/infra/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

var MeterProvider *metric.MeterProvider
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
	// Metrics
	err = initMetrics(res, ctx)
	if err != nil {
		return err
	}
	otel.SetMeterProvider(MeterProvider)
	logger.Logger.WithContext(ctx).Info("OpenTelemetry initialized successfully.")
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
	Tracer = otel.Tracer("events-service")
	return nil
}

func initMetrics(res *resource.Resource, ctx context.Context) error {
	metricExporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to create metric exporter: %w", err)
	}
	MeterProvider = metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(metricExporter)),
		metric.WithResource(res),
	)
	return nil
}

func Shutdown(ctx context.Context) error {
	if TracerProvider != nil {
		if err := TracerProvider.Shutdown(ctx); err != nil {
			logger.Logger.WithContext(ctx).WithError(err).Error("Failed to shutdown OpenTelemetry TracerProvider")
			return fmt.Errorf("failed to shutdown TracerProvider: %w", err)
		}
		logger.Logger.WithContext(ctx).Info("OpenTelemetry TracerProvider shut down.")
	}
	if MeterProvider != nil {
		if err := MeterProvider.Shutdown(ctx); err != nil {
			logger.Logger.WithContext(ctx).WithError(err).Error("Failed to shutdown OpenTelemetry MeterProvider")
			return fmt.Errorf("failed to shutdown MeterProvider: %w", err)
		}
		logger.Logger.WithContext(ctx).Info("OpenTelemetry MeterProvider shut down.")
	}
	return nil
}
