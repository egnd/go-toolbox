package victoria

import (
	vict "github.com/VictoriaMetrics/metrics"
	"github.com/egnd/go-toolbox/metrics"
)

// Histo is a histogram victoria metric.
type Histo struct {
	*vict.Histogram
	opts   *Opts
	labels metrics.Labels
}

// NewHisto is a factory method for Histogram.
func NewHisto(opts *Opts, labels ...string) *Histo {
	res := Histo{ //nolint:exhaustruct
		opts:   opts,
		labels: metrics.NewLabels(labels),
	}

	res.Histogram = vict.GetOrCreateHistogram(res.opts.ToString(&res.labels))

	return &res
}

// With updates metric labels values.
func (m *Histo) With(labelsVals ...string) metrics.Histo {
	if len(labelsVals) == 0 {
		return m
	}

	res := *m
	res.labels = m.labels.With(labelsVals...)
	res.Histogram = vict.GetOrCreateHistogram(res.opts.ToString(&res.labels))

	return &res
}
