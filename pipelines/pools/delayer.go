package pools

import (
	"context"
	"sync"
	"time"

	"github.com/egnd/go-toolbox/pipelines"
)

type delayerEvent struct {
	deadline time.Time
	callback func() error
}

// Delayer is a struct for delayed tasks execution.
type Delayer struct {
	ctx    context.Context //nolint:containedctx
	cancel context.CancelFunc
	ttl    time.Duration
	mx     sync.Mutex
	wg     sync.WaitGroup
	events map[string]delayerEvent
	stop   chan struct{}
}

// NewDelayer is a factory method for Delayer.
func NewDelayer(ctx context.Context, ttl time.Duration) *Delayer {
	var cancel context.CancelFunc

	ctx, cancel = context.WithCancel(ctx)

	return &Delayer{ //nolint:exhaustruct
		ctx:    ctx,
		cancel: cancel,
		ttl:    ttl,
		events: make(map[string]delayerEvent),
		stop:   make(chan struct{}),
	}
}

// Push is pushing a task into pool.
func (d *Delayer) Push(task pipelines.Task) error {
	if task == nil {
		return nil
	}

	d.mx.Lock()
	event, newEvent := d.events[task.ID()]
	event.deadline = time.Now().Add(d.ttl)
	event.callback = task.Do
	d.events[task.ID()] = event
	d.mx.Unlock()

	if !newEvent {
		d.wg.Add(1)

		go d.waitForEvent(task.ID())
	}

	return nil
}

// Wait blocks until tasks are completed.
func (d *Delayer) Wait() {
	d.wg.Wait()
}

// Close is stopping Delayer.
func (d *Delayer) Close() error {
	d.cancel()

	return nil
}

func (d *Delayer) waitForEvent(eventID string) {
	defer d.wg.Done()

	event := d.getEvent(eventID)

	if event.deadline.IsZero() {
		return
	}

	timer := time.NewTimer(time.Until(event.deadline))

	for {
		select {
		case <-timer.C:
			if event = d.getEvent(eventID); event.deadline.IsZero() {
				return
			}

			if event.deadline.Before(time.Now()) {
				_ = event.callback()

				d.removeEvent(eventID)

				return
			}

			timer.Reset(time.Until(event.deadline))
		case <-d.ctx.Done():
			return
		}
	}
}

func (d *Delayer) removeEvent(eventID string) {
	d.mx.Lock()
	defer d.mx.Unlock()

	delete(d.events, eventID)
}

func (d *Delayer) getEvent(eventID string) delayerEvent {
	d.mx.Lock()
	defer d.mx.Unlock()

	return d.events[eventID]
}
