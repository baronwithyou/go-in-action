package runner

import (
	"fmt"
)


// 需要实现的逻辑：接力棒到第四个人的时候，比赛结束

func Run2() {
	baton := make(chan int)

	wg.Add(1)

	go Runner2(baton)

	baton <- 1

	wg.Wait()
}

func Runner2(baton chan int) {
	defer wg.Done()

	runner := <-baton

	fmt.Printf("Runner %d running\n", runner)

	if runner != 4 {
		// 如果使用下面这种写法，会有问题，会导致runner = 4在不同的线程触发两次
		// runner++ 
		newRunner := runner + 1

		wg.Add(1)
		go Runner2(baton)

		baton <- newRunner
	}

	if runner == 4 {
		fmt.Println("Race Over")
		close(baton)
		return
	}
}