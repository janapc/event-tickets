package telemetry

import (
	"context"
	"fmt"

	"github.com/janapc/event-tickets/payments/internal/infra/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	m "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

var MeterProvider *metric.MeterProvider
var TracerProvider *sdktrace.TracerProvider
var Tracer trace.Tracer
var Meter m.Meter

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
	Tracer = otel.Tracer("tracer-payments-service")
	return nil
}

// func initTracer(res *resource.Resource) error {
// 	traceExporter, err := stdouttrace.New(
// 		stdouttrace.WithPrettyPrint())
// 	if err != nil {
// 		return err
// 	}

// 	TracerProvider = sdktrace.NewTracerProvider(
// 		sdktrace.WithBatcher(traceExporter,
// 			// Default is 5s. Set to 1s for demonstrative purposes.
// 			sdktrace.WithBatchTimeout(time.Second)),
// 		sdktrace.WithResource(res),
// 	)
// 	Tracer = otel.Tracer("tracer-payments-service")
// 	return nil
// }

func initMetrics(res *resource.Resource, ctx context.Context) error {
	metricExporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to create metric exporter: %w", err)
	}
	MeterProvider = metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(metricExporter)),
		metric.WithResource(res),
	)
	Meter = otel.Meter("metrics-payments-service")
	return nil
}

func Shutdown(ctx context.Context) {
	if TracerProvider != nil {
		if err := TracerProvider.Shutdown(ctx); err != nil {
			logger.Logger.WithContext(ctx).Error("Failed to shutdown OpenTelemetry TracerProvider")
		} else {
			logger.Logger.WithContext(ctx).Info("OpenTelemetry TracerProvider shut down.")
		}
	}
}
