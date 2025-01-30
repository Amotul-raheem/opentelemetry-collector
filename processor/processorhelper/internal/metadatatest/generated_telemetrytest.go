// Code generated by mdatagen. DO NOT EDIT.

package metadatatest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/metric/metricdata/metricdatatest"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/processortest"
)

type Telemetry struct {
	componenttest.Telemetry
}

func SetupTelemetry(opts ...componenttest.TelemetryOption) Telemetry {
	return Telemetry{Telemetry: componenttest.NewTelemetry(opts...)}
}

func (tt *Telemetry) NewSettings() processor.Settings {
	set := processortest.NewNopSettings()
	set.ID = component.NewID(component.MustNewType("processorhelper"))
	set.TelemetrySettings = tt.NewTelemetrySettings()
	return set
}

func (tt *Telemetry) AssertMetrics(t *testing.T, expected []metricdata.Metrics, opts ...metricdatatest.Option) {
	var md metricdata.ResourceMetrics
	require.NoError(t, tt.Reader.Collect(context.Background(), &md))
	// ensure all required metrics are present
	for _, want := range expected {
		got := getMetricFromResource(want.Name, md)
		metricdatatest.AssertEqual(t, want, got, opts...)
	}

	// ensure no additional metrics are emitted
	require.Equal(t, len(expected), lenMetrics(md))
}

func AssertEqualProcessorIncomingItems(t *testing.T, tt componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_processor_incoming_items",
		Description: "Number of items passed to the processor. [alpha]",
		Unit:        "{items}",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got := getMetric(t, tt, "otelcol_processor_incoming_items")
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualProcessorOutgoingItems(t *testing.T, tt componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_processor_outgoing_items",
		Description: "Number of items emitted from the processor. [alpha]",
		Unit:        "{items}",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got := getMetric(t, tt, "otelcol_processor_outgoing_items")
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func getMetric(t *testing.T, tt componenttest.Telemetry, name string) metricdata.Metrics {
	var md metricdata.ResourceMetrics
	require.NoError(t, tt.Reader.Collect(context.Background(), &md))
	return getMetricFromResource(name, md)
}

func getMetricFromResource(name string, got metricdata.ResourceMetrics) metricdata.Metrics {
	for _, sm := range got.ScopeMetrics {
		for _, m := range sm.Metrics {
			if m.Name == name {
				return m
			}
		}
	}

	return metricdata.Metrics{}
}

func lenMetrics(got metricdata.ResourceMetrics) int {
	metricsCount := 0
	for _, sm := range got.ScopeMetrics {
		metricsCount += len(sm.Metrics)
	}

	return metricsCount
}
