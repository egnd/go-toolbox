// Package metrics is a wrapper for different metrics clients.
package metrics

// Increment is an interface for increment metric.
//
//go:generate mockery --name=Increment --dir=. --output=mocks
type Increment interface {
	Inc()
	With(...string) Increment
}

// Counter is an interface for counter metric.
//
//go:generate mockery --name=Counter --dir=. --output=mocks
type Counter interface {
	Set(float64)
	With(...string) Counter
}

// Histo is an interface for histogram metric.
//
//go:generate mockery --name=Histo --dir=. --output=mocks
type Histo interface {
	Update(float64)
	With(...string) Histo
}
