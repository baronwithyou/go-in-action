package main

import (
	"fmt"
	"github.com/baronwithyou/go-in-action/chapter6/05baton/runner"
)


// 当调用同个包的方法，运行时需要编译调用的方法（go run *.go）
func main() {
	runner.Run()

	fmt.Println()

	runner.Run2()
}

