package assign_test

import (
	"fmt"
	"testing"

	"github.com/egnd/go-toolbox/pipelines/assign"

	"github.com/stretchr/testify/assert"
)

func Test_Sticky(t *testing.T) {
	cases := []struct {
		val string
		cnt uint64
		res uint64
	}{
		{"asdfsdgsg", 3, 1},
	}

	for k, test := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			assert.EqualValues(t, test.res, assign.Sticky(test.val, test.cnt))
		})
	}
}
