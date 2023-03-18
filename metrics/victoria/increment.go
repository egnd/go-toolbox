package victoria

import (
	vict "github.com/VictoriaMetrics/metrics"
	"github.com/egnd/go-toolbox/metrics"
)

type increment struct {
	*vict.FloatCounter
}

func (incr *increment) Inc() {
	incr.Add(1)
}

// Increment is a prom metric.
type Increment struct {
	builder
	opts Opts
}

// NewIncrement factory method for Increment.
func NewIncrement(opts Opts, labels ...string) *Increment {
	return &Increment{
		builder: newBuilder(labels),
		opts:    opts,
	}
}

// With append new values.
func (m *Increment) With(labelsAndValues ...string) metrics.IncrementBuilder {
	m.builder.append(labelsAndValues)

	return m
}

// Build metric instance.
func (m *Increment) Build() metrics.Increment {
	return &increment{vict.GetOrCreateFloatCounter(m.opts.ToString(m.values()))}
}
