package fb2_test

import (
	"bytes"
	"fmt"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/egnd/go-toolbox/xmlparse"
	"github.com/egnd/go-toolbox/xmlparse/fb2"
	"github.com/stretchr/testify/assert"
)

var htmlTagPattern = regexp.MustCompile("<.+?>")

func Test_Parsing(t *testing.T) {
	books := map[string][]byte{}

	items, err := filepath.Glob("examples/*.fb2")
	assert.NoError(t, err)
	for _, item := range items {
		fb2data, err := getBookData(item)
		assert.NoError(t, err)
		books[item] = fb2data
	}

	fb2data, err := getBookData("examples/small.xml")
	assert.NoError(t, err)
	books["examples/small.xml"] = fb2data

	for strPath, data := range books {
		t.Run(path.Base(strPath), func(tt *testing.T) {
			t.Log("book", path.Base(strPath))

			var marshRes fb2.File
			err = xmlparse.Unmarshal(data, &marshRes)
			assert.NoError(t, err)
			fixRes(&marshRes, len(books) == 1)

			parseRes, err := fb2.NewFile(xmlparse.NewDecoder(bytes.NewReader(data)))
			assert.NoError(t, err)
			fixRes(&parseRes, len(books) == 1)

			assert.EqualValues(t, marshRes, parseRes)
		})
	}
}

func fixRes(fb2 *fb2.File, skip bool) {
	if !skip {
		fb2.Binary = nil
	}

	for k, v := range fb2.Binary {
		fb2.Binary[k].Data = fmt.Sprint(len([]rune(strings.TrimSpace(v.Data))))
	}

	for k, v := range fb2.Description {
		for kk, vv := range v.TitleInfo {
			if !skip {
				fb2.Description[k].TitleInfo[kk].Sequence = nil
			}

			for kkk, vvv := range vv.Annotation {
				fb2.Description[k].TitleInfo[kk].Annotation[kkk].HTML = strings.TrimSpace(htmlTagPattern.ReplaceAllString(vvv.HTML, ""))
			}

			for kkk, vvv := range vv.Sequence {
				if vvv.Number == "0" {
					fb2.Description[k].TitleInfo[kk].Sequence[kkk].Number = ""
				}
			}
		}

		for kk, vv := range v.SrcTitleInfo {
			if !skip {
				fb2.Description[k].SrcTitleInfo[kk].Sequence = nil
			}

			for kkk, vvv := range vv.Annotation {
				fb2.Description[k].SrcTitleInfo[kk].Annotation[kkk].HTML = strings.TrimSpace(htmlTagPattern.ReplaceAllString(vvv.HTML, ""))
			}

			for kkk, vvv := range vv.Sequence {
				if vvv.Number == "0" {
					fb2.Description[k].SrcTitleInfo[kk].Sequence[kkk].Number = ""
				}
			}
		}

		for kk, vv := range v.PublishInfo {
			for kkk, vvv := range vv.Sequence {
				if vvv.Number == "0" {
					fb2.Description[k].PublishInfo[kk].Sequence[kkk].Number = ""
				}
			}
		}

		for kk := range v.DocInfo {
			if !skip {
				fb2.Description[k].DocInfo[kk].Publishers = nil
			}
		}
	}
}
