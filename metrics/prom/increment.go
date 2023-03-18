package prom

import (
	"github.com/egnd/go-toolbox/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Increment is a prom metric.
type Increment struct {
	builder
	factory *prometheus.CounterVec
}

// NewIncrement factory method for Increment.
func NewIncrement(opts prometheus.CounterOpts, labels ...string) *Increment {
	res := Increment{ //nolint:exhaustruct
		builder: newBuilder(labels),
	}

	res.factory = promauto.NewCounterVec(opts, res.labels)

	return &res
}

// With append new values.
func (m *Increment) With(labelsAndValues ...string) metrics.IncrementBuilder {
	m.builder.append(labelsAndValues)

	return m
}

// Build metric instance.
func (m *Increment) Build() metrics.Increment {
	return m.factory.WithLabelValues(m.values()...)
}
