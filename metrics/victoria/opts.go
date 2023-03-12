// Package victoria wraps victoria metrics.
package victoria

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/egnd/go-toolbox/metrics"
)

// Opts is a sctruct for metric name.
type Opts struct {
	Namespace string
	Subsystem string
	Name      string
}

// ToString converts Opts struct to metric name.
func (o *Opts) ToString(labels *metrics.Labels) string {
	name := o.buildName()

	if labelsStr := o.buildLabels(labels); labelsStr != "" {
		return fmt.Sprintf("%s{%s}", name, labelsStr)
	}

	return name
}

func (o *Opts) buildName() string {
	name := make([]string, 0, 3) //nolint:gomnd

	if o.Namespace != "" {
		name = append(name, o.Namespace)
	}

	if o.Subsystem != "" {
		name = append(name, o.Subsystem)
	}

	if o.Name != "" {
		name = append(name, o.Name)
	}

	return strings.Join(name, "_")
}

func (o *Opts) buildLabels(labels *metrics.Labels) string {
	names := labels.Names()
	if len(names) == 0 {
		return ""
	}

	var res bytes.Buffer

	for num, val := range labels.Values() {
		res.WriteString(fmt.Sprintf(`%s="%s",`, names[num], val))
	}

	return strings.TrimRight(res.String(), ",")
}
