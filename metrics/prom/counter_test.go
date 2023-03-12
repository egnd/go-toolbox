package prom_test

import (
	"fmt"
	"testing"

	"github.com/egnd/go-toolbox/metrics/prom"
	"github.com/prometheus/client_golang/prometheus"
)

func Test_Counter(t *testing.T) {
	for k, test := range []struct {
		labels []string
		with   []string
		val    float64
	}{
		{
			labels: []string{"label1"},
			with:   []string{"label1", "val1"},
			val:    100,
		},
		{
			val: 100,
		},
	} {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			prom.NewCounter(prometheus.GaugeOpts{Name: "counter" + fmt.Sprint(k+1)}, test.labels...).
				With(test.with...).Set(test.val)
		})
	}
}
