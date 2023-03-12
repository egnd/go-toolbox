package prom_test

import (
	"fmt"
	"testing"

	"github.com/egnd/go-toolbox/metrics/prom"
	"github.com/prometheus/client_golang/prometheus"
)

func Test_Histo(t *testing.T) {
	for k, test := range []struct {
		labels []string
		with   []string
		val    float64
	}{
		{
			labels: []string{"label1"},
			with:   []string{"label1", "val1"},
		},
		{},
	} {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			prom.NewHisto(prometheus.HistogramOpts{Name: "histo" + fmt.Sprint(k+1)}, test.labels...).
				With(test.with...).Update(test.val)
		})
	}
}
