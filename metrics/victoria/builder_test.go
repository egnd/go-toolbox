package victoria

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_builder(t *testing.T) {
	labels := []string{"param1", "param2", "param3"}
	builder := newBuilder(labels)

	for k, test := range []struct {
		lvs    []string
		values string
	}{
		{
			lvs: []string{
				"param1", "val1",
				"param3", "val3",
				"param4", "val4"},
			values: `{param1="val1",param2="",param3="val3"}`,
		},
		{
			lvs:    []string{"param2", "val2"},
			values: `{param1="val1",param2="val2",param3="val3"}`,
		},
		{
			lvs:    []string{"param24", "val24"},
			values: `{param1="val1",param2="val2",param3="val3"}`,
		},
		{
			lvs: []string{
				"param2", "val22",
				"param0", "val0",
				"param3", "val33",
			},
			values: `{param1="val1",param2="val22",param3="val33"}`,
		},
		{
			values: `{param1="val1",param2="val22",param3="val33"}`,
		},
		{
			lvs: []string{
				"param2", "val222",
				"param1",
			},
			values: `{param1="val1",param2="val222",param3="val33"}`,
		},
	} {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			builder.append(test.lvs)
			assert.EqualValues(t, test.values, builder.values())
		})
	}
}

func Test_builder_empty(t *testing.T) {
	assert.EqualValues(t, "", newBuilder([]string{}).values())
}
