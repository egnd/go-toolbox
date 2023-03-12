package victoria

import (
	vict "github.com/VictoriaMetrics/metrics"
	"github.com/egnd/go-toolbox/metrics"
)

// Increment is a victoria metric.
type Increment struct {
	*vict.Counter
	opts   *Opts
	labels metrics.Labels
}

// NewIncrement factory method for Increment.
func NewIncrement(opts *Opts, labels ...string) *Increment {
	res := Increment{ //nolint:exhaustruct
		opts:   opts,
		labels: metrics.NewLabels(labels),
	}

	res.Counter = vict.GetOrCreateCounter(res.opts.ToString(&res.labels))

	return &res
}

// With updates metric labels values.
func (m *Increment) With(labelsVals ...string) metrics.Increment {
	if len(labelsVals) == 0 {
		return m
	}

	res := *m
	res.labels = m.labels.With(labelsVals...)
	res.Counter = vict.GetOrCreateCounter(res.opts.ToString(&res.labels))

	return &res
}
