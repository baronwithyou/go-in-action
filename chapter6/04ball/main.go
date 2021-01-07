package main

import (
	"sync"
	"math/rand"
	"fmt"
)

var (
	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	court := make(chan int)

	go play("baron", court)
	go play("ling", court)

	court <- 1

	wg.Wait()
}

func play(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <- court
		if !ok {
			fmt.Printf("%s won the game\n", name)
			return
		}

		n := rand.Intn(100)
		if n % 13 == 0 {
			close(court)

			fmt.Printf("%s miss the ball\n", name)
			return
		}

		fmt.Printf("%s hit the ball %d\n", name, ball)
		ball++
		court <- ball
	}
}