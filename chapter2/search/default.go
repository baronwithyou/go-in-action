package search

/**
搜索数据用的默认匹配器
*/

type defaultMatcher struct {
}

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

func (d defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return []*Result{}, nil
}
