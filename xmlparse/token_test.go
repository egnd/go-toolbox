package xmlparse_test

import (
	"encoding/xml"
	"fmt"
	"io"
	"testing"

	"github.com/egnd/go-toolbox/xmlparse"
	"github.com/egnd/go-toolbox/xmlparse/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_TokenRead(t *testing.T) {
	cases := []struct {
		res   string
		err   error
		mocks func(*mocks.TokenReader)
	}{
		{
			err:   io.EOF,
			mocks: func(r *mocks.TokenReader) { r.On("Token").Return(nil, io.EOF) },
		},
		{
			res: "klgshjsk",
			mocks: func(r *mocks.TokenReader) {
				r.On("Token").Return(xml.CharData([]byte("klgshjsk")), nil).Once()
				r.On("Token").Return(xml.EndElement{Name: xml.Name{Local: "asdfds"}}, nil).Once()
			},
		},
	}

	for k, test := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			reader := &mocks.TokenReader{}
			test.mocks(reader)

			res, err := xmlparse.TokenRead("asdfds", reader)
			assert.EqualValues(t, test.res, res)
			assert.EqualValues(t, test.err, err)
			reader.AssertExpectations(t)
		})
	}
}

func Test_TokenSkip(t *testing.T) {
	cases := []struct {
		res   string
		err   error
		mocks func(*mocks.TokenReader)
	}{
		{
			err:   io.EOF,
			mocks: func(r *mocks.TokenReader) { r.On("Token").Return(nil, io.EOF) },
		},
		{
			mocks: func(r *mocks.TokenReader) {
				r.On("Token").Return(xml.EndElement{Name: xml.Name{Local: "asdfds"}}, nil).Once()
			},
		},
	}

	for k, test := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			reader := &mocks.TokenReader{}
			test.mocks(reader)
			err := xmlparse.TokenSkip("asdfds", reader)
			assert.EqualValues(t, test.err, err)
			reader.AssertExpectations(t)
		})
	}
}
