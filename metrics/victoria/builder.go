package victoria

import (
	"bytes"
	"fmt"
	"sort"
)

type builder struct {
	labels []string
	index  map[string]string
}

func newBuilder(labels []string) builder {
	sort.Strings(labels)

	index := make(map[string]string, len(labels))
	for _, val := range labels {
		index[val] = ""
	}

	return builder{
		labels: labels,
		index:  index,
	}
}

func (b *builder) append(lvs []string) {
	if len(lvs) == 0 {
		return
	}

	var exists bool
	for num := range lvs {
		if num%2 != 0 {
			continue
		}

		if num+1 == len(lvs) {
			break
		}

		if _, exists = b.index[lvs[num]]; exists {
			b.index[lvs[num]] = lvs[num+1]
		}
	}
}

func (b builder) values() string {
	if len(b.labels) == 0 {
		return ""
	}

	var res bytes.Buffer

	res.WriteRune('{')

	for num, label := range b.labels {
		if num > 0 {
			res.WriteRune(',')
		}

		res.WriteString(fmt.Sprintf(`%s="%s"`, label, b.index[label]))
	}

	res.WriteRune('}')

	return res.String()
}
