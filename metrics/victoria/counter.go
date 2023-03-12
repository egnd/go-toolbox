package victoria

import (
	vict "github.com/VictoriaMetrics/metrics"
	"github.com/egnd/go-toolbox/metrics"
)

// Counter is a victoria metric.
type Counter struct {
	*vict.FloatCounter
	opts   *Opts
	labels metrics.Labels
}

// NewCounter is a factory method for Counter.
func NewCounter(opts *Opts, labels ...string) *Counter {
	res := Counter{ //nolint:exhaustruct
		opts:   opts,
		labels: metrics.NewLabels(labels),
	}

	res.FloatCounter = vict.GetOrCreateFloatCounter(res.opts.ToString(&res.labels))

	return &res
}

// With updates metric labels values.
func (m *Counter) With(labelsVals ...string) metrics.Counter {
	if len(labelsVals) == 0 {
		return m
	}

	res := *m
	res.labels = m.labels.With(labelsVals...)
	res.FloatCounter = vict.GetOrCreateFloatCounter(res.opts.ToString(&res.labels))

	return &res
}
