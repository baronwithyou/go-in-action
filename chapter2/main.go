package main

import (
	"log"
	"os"

	_ "github.com/baronwithyou/go-in-action/chapter2/matchers"
	"github.com/baronwithyou/go-in-action/chapter2/search"
)

/**
1. 将main函数和每个方法的壳写好
2. 将每个文件的理解作为注释写在每个文件中
3. 开始编写代码
*/

func main() {
	log.SetOutput(os.Stderr)

	search.Run("president")
}
