package fb2_test

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/egnd/go-toolbox/xmlparse"
	"github.com/egnd/go-toolbox/xmlparse/fb2"
	"github.com/egnd/go-toolbox/xmlparse/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_NewFile(t *testing.T) {
	cases := []struct {
		res fb2.File
		err error
	}{
		{res: fb2.File{
			Description: []fb2.Description{
				{
					TitleInfo: []fb2.TitleInfo{
						{
							Genre: []string{"adventure"},
							Author: []fb2.Author{
								{
									FirstName:  []string{"name1"},
									MiddleName: []string{"patr"},
									LastName:   []string{"surname"},
									Nickname:   []string{"nickname1"},
									HomePage:   []string{"http://localhost", "http://localhost2"},
									Email:      []string{"test@test.test", "test@test.test"},
									ID:         []string{"d9f78s98s98"},
								},
							},
							BookTitle:  []string{"test book title"},
							Annotation: []fb2.Annotation{{HTML: "test annotation"}},
							Keywords:   []string{"test keywords"},
							Date:       []string{"2022-01-31"},
							Coverpage: []fb2.Cover{
								{
									Images: []fb2.Image{
										{Type: "", Href: "#mv3.jpg", Alt: "", Title: "", ID: ""},
									},
								},
							},
							Lang:    []string{"ru"},
							SrcLang: []string{"en"},
							Translator: []fb2.Author{
								{
									FirstName:  []string{"name1"},
									MiddleName: []string{"patr"},
									LastName:   []string{"surname"},
									Nickname:   []string{"nickname1"},
									HomePage:   []string{"http://localhost", "http://localhost2"},
									Email:      []string{"test@test.test", "test@test.test"},
									ID:         []string{"d9f78s98s98"},
								},
							},
							Sequence: []fb2.Sequence{
								{Name: "test seq", Number: "1"},
							},
						},
					},
					SrcTitleInfo: []fb2.TitleInfo{
						{
							Genre: []string{"adventure"},
							Author: []fb2.Author{
								{
									FirstName:  []string{"name1"},
									MiddleName: []string{"patr"},
									LastName:   []string{"surname"},
									Nickname:   []string{"nickname1"},
									HomePage:   []string{"http://localhost", "http://localhost2"},
									Email:      []string{"test@test.test", "test@test.test"},
									ID:         []string{"d9f78s98s98"},
								},
							},
							BookTitle:  []string{"test book title"},
							Annotation: []fb2.Annotation{{HTML: "testannotation"}},
							Keywords:   []string{"test keywords"},
							Date:       []string{"2022-01-31"},
							Coverpage: []fb2.Cover{
								{
									Images: []fb2.Image{
										{Type: "", Href: "#mv3.jpg", Alt: "", Title: "", ID: ""},
									},
								},
							},
							Lang:    []string{"ru"},
							SrcLang: []string{"en"},
							Translator: []fb2.Author{
								{
									FirstName:  []string{"name1"},
									MiddleName: []string{"patr"},
									LastName:   []string{"surname"},
									Nickname:   []string{"nickname1"},
									HomePage:   []string{"http://localhost", "http://localhost2"},
									Email:      []string{"test@test.test", "test@test.test"},
									ID:         []string{"d9f78s98s98"},
								},
							},
							Sequence: []fb2.Sequence{
								{Name: "test seq", Number: "1"},
							},
						},
					},
					DocInfo: []fb2.DocInfo{
						{
							Authors: []fb2.Author{
								{
									FirstName:  []string{"name1"},
									MiddleName: []string{"patr"},
									LastName:   []string{"surname"},
									Nickname:   []string{"nickname1"},
									HomePage:   []string{"http://localhost", "http://localhost2"},
									Email:      []string{"test@test.test", "test@test.test"},
									ID:         []string{"d9f78s98s98"},
								},
							},
							SrcURL:  []string{"http://localhost"},
							ID:      []string{"EC7E8CA8-9E35-48C6-A81B-3F8566FC024C"},
							Version: []string{"1.1"},
							Publishers: []fb2.Author{
								{
									FirstName:  []string{"name"},
									MiddleName: []string{"patr"},
									LastName:   []string{"surname"},
									Nickname:   []string{"nickname"},
									HomePage:   []string{"http://localhost", "http://localhost2"},
									Email:      []string{"test@test.test", "test@test.test"},
									ID:         []string{"d9f78s98s98"},
								},
							},
						},
					},
					PublishInfo: []fb2.Publisher{
						{
							BookAuthor: []string{"test author"},
							BookName:   []string{"test book"},
							Publisher:  []string{"test publisher"},
							City:       []string{"test city"},
							Year:       []string{"2022"},
							ISBN:       []string{"d90f7s7fs97s9"},
							Sequence: []fb2.Sequence{
								{Name: "test seq", Number: "1"},
							},
						},
					},
					CustomInfo: []fb2.CustomInfo{
						{InfoType: "purchased", Data: "false"},
					},
				},
			},
			Binary: []fb2.Binary{
				{ID: "mv3.jpg", ContentType: "application/octet-stream", Data: "CAZABIADASIAAhEBAxEB"},
			},
		}},
	}

	fb2data, _ := getBookData("examples/small.xml")

	for k, test := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			res, err := fb2.NewFile(xmlparse.NewDecoder(bytes.NewReader(fb2data)))
			assert.EqualValues(t, test.err, err)
			assert.EqualValues(t, test.res, res)
		})
	}
}

func Test_NewFile_Errors(t *testing.T) {
	cases := []struct {
		err   error
		mocks func(*mocks.TokenReader)
	}{
		{
			mocks: func(r *mocks.TokenReader) {
				r.On("Token").Return(nil, io.EOF)
			},
		},
		{
			err: io.EOF,
			mocks: func(r *mocks.TokenReader) {
				r.On("Token").Return(xml.StartElement{Name: xml.Name{Local: "description"}}, nil).Once()
				r.On("Token").Return(nil, io.EOF).Once()
			},
		},
	}

	for k, test := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			reader := &mocks.TokenReader{}
			test.mocks(reader)
			_, err := fb2.NewFile(reader)
			assert.EqualValues(t, test.err, err)
			reader.AssertExpectations(t)
		})
	}
}

func getBookData(path string) ([]byte, error) {
	book, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer book.Close()

	return io.ReadAll(book)
}
