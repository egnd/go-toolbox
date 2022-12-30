package xmlparse_test

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/egnd/go-toolbox/xmlparse"
	"github.com/egnd/go-toolbox/xmlparse/fb2"
)

func Benchmark_Decoders(b *testing.B) {
	fb2data, err := getBookData("fb2/examples/big.xml")
	if err != nil {
		b.Error(err)
	}

	b.Run("xml", func(bb *testing.B) {
		for i := 0; i < bb.N; i++ {
			if err := xml.NewDecoder(bytes.NewBuffer(fb2data)).Decode(&fb2.File{}); err != nil {
				bb.Error(err)
			}
		}
	})

	b.Run("xmlparse", func(bb *testing.B) {
		for i := 0; i < bb.N; i++ {
			if err := xmlparse.NewDecoder(bytes.NewBuffer(fb2data)).Decode(&fb2.File{}); err != nil {
				bb.Error(err)
			}
		}
	})
}

func Benchmark_Parsers_XML(b *testing.B) {
	table := []int{0, 1, 10, 100, 500, 1000}

	fb2data, err := getBookData("fb2/examples/big.xml")
	if err != nil {
		b.Error(err)
	}

	for _, rulesCnt := range table {

		rules := make([]xmlparse.Rule, 0, rulesCnt)
		for i := 0; i < rulesCnt; i++ {
			rules = append(rules, getBenchParsingRule)
		}

		b.Run(fmt.Sprintf("rules_%d", len(rules)), func(bb *testing.B) {
			for k := 0; k < bb.N; k++ {
				if _, err := fb2.NewFile(xml.NewDecoder(bytes.NewBuffer(fb2data)), rules...); err != nil {
					bb.Error(err)
				}
			}
		})
	}
}

func Benchmark_Parsers_XMLParsing(b *testing.B) {
	table := []int{0, 1, 10, 100, 500, 1000}

	fb2data, err := getBookData("fb2/examples/big.xml")
	if err != nil {
		b.Error(err)
	}

	for _, rulesCnt := range table {

		rules := make([]xmlparse.Rule, 0, rulesCnt)
		for i := 0; i < rulesCnt; i++ {
			rules = append(rules, getBenchParsingRule)
		}

		b.Run(fmt.Sprintf("rules_%d", len(rules)), func(bb *testing.B) {
			for k := 0; k < bb.N; k++ {
				if _, err := fb2.NewFile(xmlparse.NewDecoder(bytes.NewBuffer(fb2data)), rules...); err != nil {
					bb.Error(err)
				}
			}
		})
	}
}

func getBenchParsingRule(next xmlparse.TokenHandler) xmlparse.TokenHandler {
	return func(obj interface{}, node xml.StartElement, r xmlparse.TokenReader) (err error) {
		if _, ok := obj.(*fb2.TitleInfo); ok && node.Name.Local == "keywords" {
			obj.(*fb2.TitleInfo).Keywords = []string{"test keyword"}
		}

		return next(obj, node, r)
	}
}
