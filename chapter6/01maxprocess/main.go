package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 指定一个线程，执行两个goroutine并查看结果
	// 从这个实验可以看出只使用一个逻辑调度器的时候，主线程就是当前唯一的线程，他会阻塞所有需要运行的goroutine

	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for c := 'a'; c < 'a' + 26; c++ {
				fmt.Printf("%c ", c)
			}
			fmt.Println()
		}
	}()

	go func() {
		defer wg.Done()
		// time.Sleep(time.Second)
		for count := 0; count < 3; count++ {
			for c := 'A'; c < 'A' + 26; c++ {
				fmt.Printf("%c ", c)
			}
			fmt.Println()
		}
	}()

	fmt.Println("Waiting to finish")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}
	fmt.Println()

	wg.Wait()
}