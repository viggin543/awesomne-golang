package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool manages a set of resources that can be shared safely by multiple goroutines.
type Pool struct {
	m         sync.Mutex
	resources chan io.Closer            // stuff that can be closed
	factory   func() (io.Closer, error) // creates stuff that can be closed
	closed    bool
}

var ErrPoolClosed = errors.New("pool has been closed")

func New(allocResource func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size value too small")
	}
	return &Pool{
		factory:   allocResource,
		resources: make(chan io.Closer, size),
	}, nil
}

// why calling Acquire is safe by multiple go routines ?

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil

	default:
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}
	select {
	case p.resources <- r:
		log.Println("Release:", "In Queue")
	default:
		// get here if p.resources buffer is full
		log.Println("Release:", "Closing")
		r.Close()
	}
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		return
	}
	p.closed = true
	// Close the channel before we drain the channel of its
	// resources. If we don't do this, we will have a deadlock.
	close(p.resources)
	// why deadlock ? describe the deadlock scenario.

	for r := range p.resources {
		r.Close()
	}
}
