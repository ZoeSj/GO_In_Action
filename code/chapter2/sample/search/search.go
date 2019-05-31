package search

import (
	"log"
	"sync"
)

// A map of registered matchers for searching.
// 注册用于搜索的匹配器的映射
var matchers = make(map[string]Matcher)

// Run performs the search logic.
// run执行搜索逻辑
func Run(searchTerm string) {
	// Retrieve the list of feeds to search through.
	//获取需要搜索的数据源列表
	//：= 简化变量声明运算符，用于声明一个变量，同时给这个变量赋予初值
	feeds, err := RetrieveFeeds()

	if err != nil {
		log.Fatal(err)
	}

	//创建一个无缓冲的通道，接收匹配后的结果
	// Create an unbuffered channel to receive match results to display.
	results := make(chan *Result)

	//构造一个waitGroup，以便处理所有的数据源。
	// Setup a wait group so we can process all the feeds.
	var waitGroup sync.WaitGroup

	// Set the number of goroutines we need to wait for while
	// they process the individual feeds.
	//设置需要等待处理
	//每个数据源的goroutine的数量
	waitGroup.Add(len(feeds))

	// Launch a goroutine for each feed to find the results.
	//为每个数据源启动一个goroutine来查找结果
	//这里的下划线是占位符
	for _, feed := range feeds {
		// Retrieve a matcher for the search.
		//获取一个匹配器用于查找
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// Launch the goroutine to perform the search.
		//启动一个goroutine来执行搜索
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// Launch a goroutine to monitor when all the work is done.
	//启动一个goroutine来监控是否所有的工作都做完了
	go func() {
		// Wait for everything to be processed.
		//等待所有任务完成
		waitGroup.Wait()

		// Close the channel to signal to the Display
		// function that we can exit the program.
		//用关闭通道的方式，通知display函数
		//可以退出程序了
		close(results)
	}()

	// Start displaying results as they are available and
	// return after the final result is displayed.
	//启动函数，显示返回的结果，并且，在最后一个结果显示完之后返回。
	Display(results)
}

// Register is called to register a matcher for use by the program.
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
