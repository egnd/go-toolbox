package concurr

type Pipeline interface {
	Push(Task) error
	Wait() error
	Close() error
}
