package victoria_test

import (
	"fmt"
	"testing"

	"github.com/egnd/go-toolbox/metrics"
	"github.com/egnd/go-toolbox/metrics/victoria"
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
			val:    123,
		},
	} {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			var obj metrics.HistoBuilder = victoria.NewHisto(victoria.Opts{Name: "histo" + fmt.Sprint(k+1)}, test.labels...)
			obj.With(test.with...).Build().Update(test.val)
		})
	}
}
