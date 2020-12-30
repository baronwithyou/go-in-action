package search

import (
	"fmt"
	"log"
	"sync"
)

/**
执行搜索的主控制逻辑
*/

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	results := make(chan *Result)

	// 每个feed启动一个goroutine
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		go func(matcher2 Matcher, feed2 *Feed) {
			defer wg.Done()

			result, _ := matcher.Search(feed2, searchTerm)

			for _, item := range result {
				results <- item
			}
		}(matcher, feed)
	}

	go func() {
		defer close(results)

		wg.Wait()
	}()

	// 最后通过Display打印出来
	Display(results)
}

func Display(results chan *Result) {
	for result := range results {
		fmt.Printf("Field: %v, Content: %v\n", result.Field, result.Content)
	}
}

func Register(name string, matcher Matcher) {
	matchers[name] = matcher
}
