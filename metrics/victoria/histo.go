package victoria

import (
	vict "github.com/VictoriaMetrics/metrics"
	"github.com/egnd/go-toolbox/metrics"
)

// Histo is a prom metric.
type Histo struct {
	builder
	opts Opts
}

// NewHisto factory method for Histo.
func NewHisto(opts Opts, labels ...string) *Histo {
	return &Histo{
		builder: newBuilder(labels),
		opts:    opts,
	}
}

// With append new values.
func (m *Histo) With(labelsAndValues ...string) metrics.HistoBuilder {
	m.builder.append(labelsAndValues)

	return m
}

// Build metric instance.
func (m *Histo) Build() metrics.Histo {
	return vict.GetOrCreateHistogram(m.opts.ToString(m.values()))
}
