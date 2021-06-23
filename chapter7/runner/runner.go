package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"
)

/**
signal.Stop() // 停止接收后续的任何信号
signal.Notify() // 接收中断信号
*/

type Task func(int)

type Runner struct {
	interrupt chan os.Signal

	complete chan error

	timeout <-chan time.Time

	tasks []Task
}

var ErrTimeout = errors.New("received timeout")

var ErrInterrupt = errors.New("received interrupt")

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

func (r *Runner) Add(t ...Task) {
	r.tasks = append(r.tasks, t...)
}

// 这里使用值接收者还是指针接收者???
func (r *Runner) Start() error {
	fmt.Println("Start the Runner")
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		task(id)
	}

	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		return true
	default:
		return false
	}
}
