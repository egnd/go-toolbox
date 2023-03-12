// Package prom wraps prometheus metrics.
package prom

import (
	"github.com/egnd/go-toolbox/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Counter is a prom metric.
type Counter struct {
	prometheus.Gauge
	labels  metrics.Labels
	factory *prometheus.GaugeVec
}

// NewCounter is a factory method for Counter.
func NewCounter(opts prometheus.GaugeOpts, labels ...string) *Counter {
	res := Counter{ //nolint:exhaustruct
		labels: metrics.NewLabels(labels),
	}

	res.factory = promauto.NewGaugeVec(opts, res.labels.Names())
	res.Gauge = res.factory.WithLabelValues(res.labels.Values()...)

	return &res
}

// With updates metric labels values.
func (m *Counter) With(labelsVals ...string) metrics.Counter {
	if len(labelsVals) == 0 {
		return m
	}

	res := *m
	res.labels = m.labels.With(labelsVals...)
	res.Gauge = res.factory.WithLabelValues(res.labels.Values()...)

	return &res
}
