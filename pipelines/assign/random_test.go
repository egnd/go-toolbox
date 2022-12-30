package assign_test

import (
	"fmt"
	"testing"

	"github.com/egnd/go-toolbox/pipelines/assign"

	"github.com/stretchr/testify/assert"
)

func Test_Random(t *testing.T) {
	cases := []uint64{1, 2, 3, 4, 5, 10, 50, 100}

	for k, cnt := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			if cnt == 1 {
				assert.EqualValues(t, 0, assign.Random("", cnt))
			} else {
				res := assign.Random("", cnt)
				assert.Less(t, res, cnt)
				assert.GreaterOrEqual(t, res, uint64(0))
			}
		})
	}
}
