package search

/**
用于支持不同匹配器的接口
 */

type Result struct {
	Field string
	Content string
}

type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}
