package metrics_test

import (
	"fmt"
	"testing"

	"github.com/egnd/go-toolbox/metrics"
	"github.com/stretchr/testify/assert"
)

func Test_Labels(t *testing.T) {
	names := []string{"param1", "param2", "param3"}
	labels := metrics.NewLabels(names)

	for k, test := range []struct {
		lvs    []string
		values []string
	}{
		{
			lvs: []string{
				"param1", "val1",
				"param3", "val3",
				"param4", "val4"},
			values: []string{"val1", "", "val3"},
		},
		{
			lvs:    []string{"param2", "val2"},
			values: []string{"val1", "val2", "val3"},
		},
		{
			lvs:    []string{"param24", "val24"},
			values: []string{"val1", "val2", "val3"},
		},
		{
			lvs: []string{
				"param2", "val22",
				"param0", "val0",
				"param3", "val33",
			},
			values: []string{"val1", "val22", "val33"},
		},
		{
			values: []string{"val1", "val22", "val33"},
		},
		{
			lvs: []string{
				"param2", "val222",
				"param1",
			},
			values: []string{"val1", "val222", "val33"},
		},
	} {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			labels = labels.With(test.lvs...)
			assert.EqualValues(t, names, labels.Names())
			assert.EqualValues(t, test.values, labels.Values())
		})
	}
}
