package fb2_test

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/egnd/go-toolbox/xmlparse/fb2"
	"github.com/stretchr/testify/assert"
)

func Test_NewSequence(t *testing.T) {
	cases := []struct {
		token xml.StartElement
		res   fb2.Sequence
		err   error
	}{
		{
			token: xml.StartElement{Attr: []xml.Attr{
				{Name: xml.Name{Local: "name"}, Value: "test-name"},
				{Name: xml.Name{Local: "number"}, Value: "asdf"},
			}},
			res: fb2.Sequence{Name: "test-name"},
		},
	}

	for k, test := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			res, err := fb2.NewSequence(test.token)
			assert.EqualValues(t, test.res, res)
			assert.EqualValues(t, test.err, err)
		})
	}
}

func Test_Sequence_String(t *testing.T) {
	cases := []struct {
		item fb2.Sequence
		res  string
	}{
		{
			res:  "name (num)",
			item: fb2.Sequence{Number: "num", Name: "name"},
		},
		{
			res:  "name (num)",
			item: fb2.Sequence{Number: "num", Name: "name,"},
		},
		{
			item: fb2.Sequence{Name: ","},
		},
		{
			res: "",
		},
	}
	for k, test := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			assert.EqualValues(t, test.res, test.item.String())
		})
	}
}
