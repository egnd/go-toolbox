// Package metrics is a wrapper for different metrics clients.
package metrics

//go:generate mockery --name=Increment --dir=. --output=mocks

// Increment is an interface for increment metric.
type Increment interface {
	Inc()
	Add(float64)
}

type IncrementBuilder interface {
	With(...string) IncrementBuilder
	Build() Increment
}

// Counter is an interface for counter metric.
type Counter interface {
	Set(float64)
}

type CounterBuilder interface {
	With(...string) CounterBuilder
	Build() Counter
}

// Histo is an interface for histogram metric.
type Histo interface {
	Update(float64)
}

type HistoBuilder interface {
	With(...string) HistoBuilder
	Build() Histo
}
