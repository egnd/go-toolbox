package xmlparse_test

import (
	"io"
	"os"
	"testing"

	"github.com/egnd/go-toolbox/xmlparse"
	"github.com/egnd/go-toolbox/xmlparse/fb2"
	"github.com/stretchr/testify/assert"
)

func Test_Decoder(t *testing.T) {
	fb2data, err := getBookData("fb2/examples/small.xml")
	assert.NoError(t, err)
	assert.NoError(t, xmlparse.Unmarshal(fb2data, &fb2.File{}))
}

func Test_Decoder_Error(t *testing.T) {
	assert.Error(t, xmlparse.Unmarshal([]byte{}, &fb2.File{}))
}

func getBookData(path string) ([]byte, error) {
	book, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer book.Close()

	return io.ReadAll(book)
}
