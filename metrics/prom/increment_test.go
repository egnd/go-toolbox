package prom_test

import (
	"fmt"
	"testing"

	"github.com/egnd/go-toolbox/metrics/prom"
	"github.com/prometheus/client_golang/prometheus"
)

func Test_Increment(t *testing.T) {
	for k, test := range []struct {
		labels []string
		with   []string
	}{
		{
			labels: []string{"label1"},
			with:   []string{"label1", "val1"},
		},
		{},
	} {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			prom.NewIncrement(prometheus.CounterOpts{Name: "incr" + fmt.Sprint(k+1)}, test.labels...).
				With(test.with...).Inc()
		})
	}
}
