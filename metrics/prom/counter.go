// Package prom wraps prometheus metrics.
package prom

import (
	"github.com/egnd/go-toolbox/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Counter is a prom metric.
type Counter struct {
	builder
	factory *prometheus.GaugeVec
}

// NewCounter factory method for Counter.
func NewCounter(opts prometheus.GaugeOpts, labels ...string) *Counter {
	res := Counter{ //nolint:exhaustruct
		builder: newBuilder(labels),
	}

	res.factory = promauto.NewGaugeVec(opts, res.labels)

	return &res
}

// With append new values.
func (m *Counter) With(labelsAndValues ...string) metrics.CounterBuilder {
	m.builder.append(labelsAndValues)

	return m
}

// Build metric instance.
func (m *Counter) Build() metrics.Counter {
	return m.factory.WithLabelValues(m.values()...)
}
