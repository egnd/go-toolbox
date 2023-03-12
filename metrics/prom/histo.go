package prom

import (
	"github.com/egnd/go-toolbox/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Histo is a histogram prom metric.
type Histo struct {
	prometheus.Observer
	labels  metrics.Labels
	factory *prometheus.HistogramVec
}

// NewHisto is a factory method for Histogram.
func NewHisto(opts prometheus.HistogramOpts, labels ...string) *Histo {
	res := Histo{ //nolint:exhaustruct
		labels: metrics.NewLabels(labels),
	}

	res.factory = promauto.NewHistogramVec(opts, res.labels.Names())
	res.Observer = res.factory.WithLabelValues(res.labels.Values()...)

	return &res
}

// Update histogram value.
func (m *Histo) Update(val float64) {
	m.Observe(val)
}

// With updates metric labels values.
func (m *Histo) With(labelsVals ...string) metrics.Histo {
	if len(labelsVals) == 0 {
		return m
	}

	res := *m
	res.labels = m.labels.With(labelsVals...)
	res.Observer = res.factory.WithLabelValues(res.labels.Values()...)

	return &res
}
