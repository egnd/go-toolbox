// Package victoria wraps victoria metrics.
package victoria

import (
	"fmt"
	"strings"
)

// Opts is a sctruct for metric name.
type Opts struct {
	Namespace string
	Subsystem string
	Name      string
}

// ToString converts Opts struct to metric name.
func (o *Opts) ToString(labels string) string {
	return fmt.Sprintf("%s%s", o.buildName(), labels)
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
