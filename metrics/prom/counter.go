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
	res := Counter{
		builder: newBuilder(labels),
	}

	res.factory = promauto.NewGaugeVec(opts, res.labels)

	return &res
}

func (m *Counter) With(labelsAndValues ...string) metrics.CounterBuilder {
	m.builder.append(labelsAndValues)
	return m
}

func (m *Counter) Build() metrics.Counter {
	return m.factory.WithLabelValues(m.values()...)
}
