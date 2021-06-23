package main

import (
	"fmt"
	"time"
	//runner "github.com/baronwithyou/go-in-action/chapter7/runner"
)

var (
	d = time.Second * 4
)

func main() {
	r := New(d)

	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("All task done!")
}

func task(i int) {
	time.Sleep(time.Duration(i) * time.Second)

	fmt.Printf("task %d done\n", i)
}

func createTask() Task {
	return task
}
