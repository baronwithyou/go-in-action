package main


import (
	"sync"
	"runtime"
	"fmt"
)

var wg sync.WaitGroup

func main() {
	// 设置一个线程，运行两个goroutine

	fmt.Println(runtime.NumCPU())
	return
	runtime.GOMAXPROCS(1)


	wg.Add(2)

	go printPrime("A")
	go printPrime("B")

	wg.Wait()
}

func printPrime(prefix string) {
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
				// break
			}
		}

		fmt.Printf("%s: %d\n", prefix, outer)
	}

	fmt.Println("Completed:", prefix)
}