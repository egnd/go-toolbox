package xmlparse_test

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/egnd/go-toolbox/xmlparse"
	"github.com/stretchr/testify/assert"
)

func Test_WrapRules(t *testing.T) {
	cases := []string{
		"hello world",
		"",
	}

	for k, phrase := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			var res string
			rules := []xmlparse.Rule{}
			if len(phrase) > 0 {
				for i := 0; i < len(phrase)-1; i++ {
					i := i
					rules = append(rules, func(next xmlparse.TokenHandler) xmlparse.TokenHandler {
						return func(section interface{}, node xml.StartElement, r xmlparse.TokenReader) error {
							res += phrase[i : i+1]
							return next(section, node, r)
						}
					})
				}
			}

			xmlparse.WrapRules(rules, func(_ interface{}, _ xml.StartElement, _ xmlparse.TokenReader) error {
				if len(phrase) > 0 {
					res += phrase[len(phrase)-1:]
				}

				return nil
			})(nil, xml.StartElement{}, nil)

			assert.EqualValues(t, phrase, res+"")
		})
	}
}
