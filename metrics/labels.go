package metrics

import "sort"

// Labels is a struct for Labels managing.
type Labels struct {
	names  []string
	values []string
}

// NewLabels factory method for Labels struct.
func NewLabels(names []string) Labels {
	sort.Strings(names)

	return Labels{
		names:  names,
		values: make([]string, len(names)),
	}
}

// Names returns slice of labels names.
func (l *Labels) Names() []string {
	return l.names
}

// Values returns slice of labels values.
func (l *Labels) Values() []string {
	return l.values
}

// With returns copy of Labels struct with new labels values.
func (l Labels) With(labelsAndValues ...string) Labels {
	if len(labelsAndValues) == 0 {
		return l
	}

	newVals := make(map[string]string, len(labelsAndValues)/2) //nolint:gomnd

	for num := range labelsAndValues {
		if num%2 != 0 {
			continue
		}

		if num+1 == len(labelsAndValues) {
			break
		}

		newVals[labelsAndValues[num]] = labelsAndValues[num+1]
	}

	values := make([]string, 0, len(l.names))

	for num, name := range l.names {
		if _, ok := newVals[name]; ok {
			values = append(values, newVals[name])
		} else {
			values = append(values, l.values[num])
		}
	}

	l.values = values

	return l
}
