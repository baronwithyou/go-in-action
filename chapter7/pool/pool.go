package main

import (
	"errors"
	"io"
)

/**
拥有的功能：
1. Acquire ->io.Closer
2. Release
3. Close
*/

type Pool struct {
	closed    bool
	resources chan io.Closer
	factory   Factory
}

type Factory func() (io.Closer, error)

var PoolClosedErr = errors.New("the pool has been closed")
var ResourceLackErr = errors.New("the pool is lack of resource")

// New ...
func New(fn Factory, size int) (*Pool, error) {
	return &Pool{
		factory:   fn,
		closed:    false,
		resources: make(chan io.Closer, size),
	}, nil
}

// Acquire ...
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case resource, ok := <-p.resources:
		if !ok {
			return nil, PoolClosedErr
		}
		return resource, nil
	default:
		return p.factory()
	}
}

// Release ...
func (p *Pool) Release(resource io.Closer) {
	if p.closed {
		resource.Close()
		return
	}

	select {
	case p.resources <- resource:
	default:
		resource.Close()
	}
}

// Close ...
func (p *Pool) Close() {
	if p.closed {
		return
	}

	// 将channel关了
	close(p.resources)

	// 再遍历resources关闭
	for r := range p.resources {
		r.Close()
	}

	p.closed = true
}
