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

func Test_NewPublisher(t *testing.T) {
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
				r.On("Token").Return(xml.StartElement{Name: xml.Name{Local: "book-author"}}, nil).Once()
				r.On("Token").Return(nil, io.EOF).Once()
			},
		},
	}

	for k, test := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			reader := &mocks.TokenReader{}
			test.mocks(reader)
			_, err := fb2.NewPublisher("asd", reader, nil)
			assert.EqualValues(t, test.err, err)
			reader.AssertExpectations(t)
		})
	}
}

func Test_Publisher_String(t *testing.T) {
	cases := []struct {
		item fb2.Publisher
		res  string
	}{
		{
			res: "publ (city)",
			item: fb2.Publisher{
				Publisher: []string{"", "publ"},
				City:      []string{"city"},
			},
		},
		{
			res: "",
		},
		{
			res: "city",
			item: fb2.Publisher{
				City: []string{"city"},
			},
		},
	}
	for k, test := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			assert.EqualValues(t, test.res, test.item.String())
		})
	}
}
