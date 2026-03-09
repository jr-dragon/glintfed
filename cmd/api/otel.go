package main

import (
	"context"
	"errors"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.39.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"glintfed.org/internal/data"
)

func setupOTelSDK(ctx context.Context, cfg data.Config) (cleanup func(context.Context) error, err error) {
	var cleanups []func(context.Context) error
	cleanup = func(ctx context.Context) error {
		var err error
		for _, fn := range cleanups {
			err = errors.Join(err, fn(ctx))
		}
		cleanups = nil
		return err
	}

	handleErr := func(inErr error) {
		err = errors.Join(inErr, cleanup(ctx))
	}

	otel.SetTextMapPropagator(newPropagator())

	res, err := resource.New(ctx, resource.WithAttributes(semconv.ServiceName(cfg.App.Name)))
	if err != nil {
		return nil, err
	}

	if cfg.Service.OpenTelemetry.TracingEnabled {
		tp, err := newTracerProvider(ctx, cfg.Service.OpenTelemetry, res)
		if err != nil {
			handleErr(err)
			return nil, err
		}
		cleanups = append(cleanups, tp.Shutdown)
		otel.SetTracerProvider(tp)
	}

	return
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func newTracerProvider(ctx context.Context, cfg data.OpenTelemetryConfig, res *resource.Resource) (*trace.TracerProvider, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	conn, err := grpc.NewClient(cfg.TracingEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, err
	}

	return trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithResource(res),
		trace.WithSpanProcessor(trace.NewBatchSpanProcessor(exporter)),
	), nil
}
