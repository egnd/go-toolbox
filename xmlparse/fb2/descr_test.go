package fb2_test

import (
	"encoding/xml"
	"fmt"
	"io"
	"testing"

	"github.com/egnd/go-toolbox/xmlparse/fb2"
	"github.com/egnd/go-toolbox/xmlparse/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_NewDescription(t *testing.T) {
	cases := []struct {
		err   error
		token xml.StartElement
		mocks func(*mocks.TokenReader)
	}{
		{
			err: io.EOF,
			mocks: func(r *mocks.TokenReader) {
				r.On("Token").Return(nil, io.EOF)
			},
		},
		{
			err: io.EOF,
			mocks: func(r *mocks.TokenReader) {
				r.On("Token").Return(xml.StartElement{Name: xml.Name{Local: "title-info"}}, nil).Once()
				r.On("Token").Return(nil, io.EOF).Once()
			},
		},
	}

	for k, test := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			reader := &mocks.TokenReader{}
			test.mocks(reader)
			_, err := fb2.NewDescription("asd", reader, nil)
			assert.EqualValues(t, test.err, err)
			reader.AssertExpectations(t)
		})
	}
}
