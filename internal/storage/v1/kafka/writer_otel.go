// Copyright (c) 2018 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package kafka

import (
	"context"
	"github.com/jaegertracing/jaeger/internal/storage/v2/v1adapter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"

	"github.com/jaegertracing/jaeger-idl/model/v1"
	"github.com/jaegertracing/jaeger/pkg/metrics"
	"go.uber.org/zap"
)

// NewOtelSpanWriter initiates and returns a new kafka spanwriter
func NewOtelSpanWriter(
	producer exporter.Traces,
	topic string,
	factory metrics.Factory,
	logger *zap.Logger,
) *OTelSpanWriter {
	writeMetrics := spanWriterMetrics{
		SpansWrittenSuccess: factory.Counter(metrics.Options{Name: "kafka_spans_written", Tags: map[string]string{"status": "success"}}),
		SpansWrittenFailure: factory.Counter(metrics.Options{Name: "kafka_spans_written", Tags: map[string]string{"status": "failure"}}),
	}

	exporterFactory := kafkaexporter.NewFactory()
	cfg := exporterFactory.CreateDefaultConfig().(*kafkaexporter.Config)

	// Initialize the exporter
	otelExporter, err := exporterFactory.CreateTraces(context.Background(),
		exporter.Settings{
			ID:                component.ID{},
			TelemetrySettings: component.TelemetrySettings{},
			BuildInfo:         component.BuildInfo{},
		}, cfg)
	if err != nil {
		logger.Fatal("Failed to create OTEL Kafka exporter", zap.Error(err))
	}

	return &OTelSpanWriter{
		exporter: nil,
		metrics:  writeMetrics,
	}
}

type OTelSpanWriter struct {
	exporter exporter.Traces
	metrics  spanWriterMetrics
}

// WriteSpan writes the span to kafka.
func (w *OTelSpanWriter) WriteSpan(ctx context.Context, span *model.Span) error {
	otelTraces := v1adapter.V1BatchesToTraces([]*model.Batch{{Spans: []*model.Span{span}}})
	err := w.exporter.ConsumeTraces(ctx, otelTraces)
	if err != nil {
		w.metrics.SpansWrittenFailure.Inc(1)
		return err
	}

	// Increment success metric
	w.metrics.SpansWrittenSuccess.Inc(1)
	return nil
}

// Close closes SpanWriter by closing producer
func (w *OTelSpanWriter) Close() error {
	return w.exporter.Shutdown(context.Background())
}
