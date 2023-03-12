package victoria_test

import (
	"fmt"
	"testing"

	"github.com/egnd/go-toolbox/metrics"
	"github.com/egnd/go-toolbox/metrics/victoria"
	"github.com/stretchr/testify/assert"
)

func Test_Opts(t *testing.T) {
	for k, test := range []struct {
		opts   victoria.Opts
		labels metrics.Labels
		res    string
	}{
		{
			opts:   victoria.Opts{Namespace: "ns", Subsystem: "sys", Name: "name"},
			labels: metrics.NewLabels([]string{"label1"}).With("label1", "val1"),
			res:    `ns_sys_name{label1="val1"}`,
		},
	} {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			assert.EqualValues(t, test.res, test.opts.ToString(&test.labels))
		})
	}
}
