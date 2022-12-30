package fb2_test

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/egnd/go-toolbox/xmlparse/fb2"
	"github.com/stretchr/testify/assert"
)

func Test_NewImage(t *testing.T) {
	cases := []struct {
		token xml.StartElement
		res   fb2.Image
	}{
		{
			token: xml.StartElement{Attr: []xml.Attr{
				{Name: xml.Name{Local: "type"}, Value: "111"},
				{Name: xml.Name{Local: "href"}, Value: "222"},
				{Name: xml.Name{Local: "alt"}, Value: "333"},
				{Name: xml.Name{Local: "title"}, Value: "444"},
				{Name: xml.Name{Local: "id"}, Value: "555"},
			}},
			res: fb2.Image{
				Type:  "111",
				Href:  "222",
				Alt:   "333",
				Title: "444",
				ID:    "555",
			},
		},
	}

	for k, test := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			assert.EqualValues(t, test.res, fb2.NewImage(test.token))
		})
	}
}
