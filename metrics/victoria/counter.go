package victoria

import (
	vict "github.com/VictoriaMetrics/metrics"
	"github.com/egnd/go-toolbox/metrics"
)

// Counter is a prom metric.
type Counter struct {
	builder
	opts Opts
}

// NewCounter factory method for Counter.
func NewCounter(opts Opts, labels ...string) *Counter {
	return &Counter{
		builder: newBuilder(labels),
		opts:    opts,
	}
}

// With append new values.
func (m *Counter) With(labelsAndValues ...string) metrics.CounterBuilder {
	m.builder.append(labelsAndValues)

	return m
}

// Build metric instance.
func (m *Counter) Build() metrics.Counter {
	return vict.GetOrCreateFloatCounter(m.opts.ToString(m.values()))
}
