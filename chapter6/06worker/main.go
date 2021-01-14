package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg        sync.WaitGroup
	taskNum   = 10
	workerNum = 4
)

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

		time.Sleep(time.Millisecond * 200)

		fmt.Printf("Worker %d done task %d\n", id, task)
	}
}
