package main

import (
	"sync"
	"fmt"
	"runtime"
)

var (
	counter int
	wg sync.WaitGroup
	mutex sync.Mutex
)

func main() {
	wg.Add(2)

	go incCounter("A")
	go incCounter("B")


	wg.Wait()

	fmt.Println(counter)
}

func incCounter(s string) {
	defer wg.Done()

	for i := 0; i < 2;i ++ {
		mutex.Lock()

		value := counter

		runtime.Gosched()

		value++
		counter = value

		mutex.Unlock()

		fmt.Println(s)
	}
}