package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	wg        sync.WaitGroup
	taskNum   = 10
	workerNum = 4
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	tasks := make(chan int, taskNum)

	wg.Add(workerNum)

	for i := 1; i <= workerNum; i++ {
		go worker(tasks, i)
	}

	for i := 1; i <= taskNum; i++ {
		tasks <- i
	}

	// 关闭了通道但是还是可以从通道中获取数据
	close(tasks)

	wg.Wait()
}

func worker(tasks chan int, id int) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker %d done all the tasks\n", id)
			return
		}

		time.Sleep(time.Millisecond * time.Duration(rand.Int63n(100)))

		fmt.Printf("Worker %d done task %d\n", id, task)
	}
}
