package xmlparse_test

import (
	"fmt"
	"testing"

	"github.com/egnd/go-toolbox/xmlparse"
	"github.com/stretchr/testify/assert"
)

func Test_GetStrFrom(t *testing.T) {
	cases := []struct {
		haystack []string
		needle   string
	}{
		{needle: "asdf", haystack: []string{"", "asdf", "", "hgf"}},
		{needle: "", haystack: []string{""}},
		{needle: "", haystack: []string{}},
		{needle: "", haystack: nil},
	}

	for k, test := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			assert.EqualValues(t, test.needle, xmlparse.GetStrFrom(test.haystack))
		})
	}
}
