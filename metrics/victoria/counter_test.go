package victoria_test

import (
	"fmt"
	"testing"

	"github.com/egnd/go-toolbox/metrics/victoria"
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
			victoria.NewCounter(&victoria.Opts{Name: "counter" + fmt.Sprint(k+1)}, test.labels...).
				With(test.with...).Set(test.val)
		})
	}
}
