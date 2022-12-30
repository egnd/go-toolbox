package fb2_test

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/egnd/go-toolbox/xmlparse/fb2"
	"github.com/egnd/go-toolbox/xmlparse/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_NewBody(t *testing.T) {
	cases := []struct {
		res   fb2.Body
		err   error
		mocks func(*mocks.TokenReader)
	}{
		{
			res: fb2.Body{
				HTML: "asdg",
			},
			mocks: func(r *mocks.TokenReader) {
				r.On("Token").Return(xml.CharData("asdg"), nil).Once()
				r.On("Token").Return(xml.EndElement{Name: xml.Name{Local: "asdfds"}}, nil).Once()
			},
		},
	}

	for k, test := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			reader := &mocks.TokenReader{}
			test.mocks(reader)
			res, err := fb2.NewBody("asdfds", reader)
			assert.EqualValues(t, test.res, res)
			assert.EqualValues(t, test.err, err)
			reader.AssertExpectations(t)
		})
	}
}
