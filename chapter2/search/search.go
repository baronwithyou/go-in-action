package search

/**
执行搜索的主控制逻辑
 */

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {

	// 每个matcher启动一个goroutine

	// 最后通过Display打印出来
}

func Display(results chan *Result) {

}