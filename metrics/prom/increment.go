package prom

import (
	"github.com/egnd/go-toolbox/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Increment is a prom metric.
type Increment struct {
	prometheus.Counter
	labels  metrics.Labels
	factory *prometheus.CounterVec
}

// NewIncrement factory method for Increment.
func NewIncrement(opts prometheus.CounterOpts, labels ...string) *Increment {
	res := Increment{ //nolint:exhaustruct
		labels: metrics.NewLabels(labels),
	}

	res.factory = promauto.NewCounterVec(opts, res.labels.Names())
	res.Counter = res.factory.WithLabelValues(res.labels.Values()...)

	return &res
}

// Add value to increment.
func (m *Increment) Add(val int) {
	m.Counter.Add(float64(val))
}

// With updates metric labels values.
func (m *Increment) With(labelsVals ...string) metrics.Increment {
	if len(labelsVals) == 0 {
		return m
	}

	res := *m
	res.labels = m.labels.With(labelsVals...)
	res.Counter = res.factory.WithLabelValues(res.labels.Values()...)

	return &res
}
