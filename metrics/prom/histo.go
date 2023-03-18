package prom

import (
	"github.com/egnd/go-toolbox/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type histogram struct {
	prometheus.Observer
}

func (m *histogram) Update(val float64) {
	m.Observe(val)
}

// Histo is a prom metric.
type Histo struct {
	builder
	factory *prometheus.HistogramVec
}

// NewHisto factory method for Histo.
func NewHisto(opts prometheus.HistogramOpts, labels ...string) *Histo {
	res := Histo{ //nolint:exhaustruct
		builder: newBuilder(labels),
	}

	res.factory = promauto.NewHistogramVec(opts, res.labels)

	return &res
}

// With append new values.
func (m *Histo) With(labelsAndValues ...string) metrics.HistoBuilder {
	m.builder.append(labelsAndValues)

	return m
}

// Build metric instance.
func (m *Histo) Build() metrics.Histo {
	return &histogram{m.factory.WithLabelValues(m.values()...)}
}
