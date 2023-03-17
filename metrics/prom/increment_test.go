package prom_test

import (
	"fmt"
	"testing"

	"github.com/egnd/go-toolbox/metrics"
	"github.com/egnd/go-toolbox/metrics/prom"
	"github.com/prometheus/client_golang/prometheus"
)

func Test_Increment(t *testing.T) {
	for k, test := range []struct {
		labels []string
		with   []string
		val    float64
	}{
		{
			labels: []string{"label1"},
			with:   []string{"label1", "val1"},
			val:    123,
		},
	} {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			var obj metrics.IncrementBuilder = prom.NewIncrement(prometheus.CounterOpts{Name: "incr" + fmt.Sprint(k+1)}, test.labels...)

			metric := obj.With(test.with...).Build()
			metric.Inc()
			metric.Add(test.val)
		})
	}
}
