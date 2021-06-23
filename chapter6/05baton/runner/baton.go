package runner

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
)

func Run() {

	baton := make(chan int)

	wg.Add(1)

	go Runner(baton)

	baton <- 1

	wg.Wait()
}

func Runner(baton chan int) {
	defer wg.Done()

	runner := <- baton

	fmt.Printf("Runner %d running\n", runner)

	if runner == 4 {
		fmt.Printf("Runner %d finished, Race Over\n", runner)
		close(baton)
		return
	}

	if runner != 4 {
		// 启动线程必须在写入通道之前，因为该通道是没有缓冲区的
		wg.Add(1)
		go Runner(baton)

		runner++
		baton<-runner
	}
}