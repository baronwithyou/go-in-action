package main

import (
	"fmt"
	"io"
)

type NewCloser struct {
}

func (n *NewCloser) Close() error {
	fmt.Println("I am closed !")
	return nil
}

func main() {
	pool, _ := New(fn, 2)

	r1, _ := pool.Acquire()
	r2, _ := pool.Acquire()

	pool.Release(r1)
	pool.Release(r2)

	pool.Close()

	r3, err := pool.Acquire()
	if err == nil {
		pool.Release(r3)
	}
}

func fn() (io.Closer, error) {
	return &NewCloser{}, nil
}
