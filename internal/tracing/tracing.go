package tracing

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
)

func createTracingProvider(url string) (*traceSdk.TracerProvider, error){
	telemetryExporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tracingProvider := traceSdk.NewTracerProvider(
		traceSdk.WithBatcher(telemetryExporter), // disini data akan dikirim secara batch. dikumoulkan ke jaeger 
		traceSdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("my-pos"),
			attribute.String("environment", "development"),
		)),
	)
	return tracingProvider, nil	
}

func Init(url string) error {
	tracingProvider, err := createTracingProvider(url)
	if err != nil {
		return err
	}

	otel.SetTracerProvider(tracingProvider)
	return nil
}

// Kita akan buat wrapper
func CreateSpanWrapper(ctx context.Context, name string) (context.Context, trace.Span) {
	if ctx == nil {
		ctx = context.Background()
	}

	tracerTracing := otel.Tracer(name)
	ctx, span := tracerTracing.Start(ctx, name)

	return ctx, span

}
