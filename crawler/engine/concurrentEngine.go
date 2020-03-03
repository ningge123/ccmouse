package engine

import (
	"log"
	"ccmouse/crawler/fetcher"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkNum   int
	ItemChan  chan Item
}

type Scheduler interface {
	Submit(Request)
	ReadyNotifier
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

// 全局变量集合去重url

var uniqueUrl = make(map[string]bool)

func (concurrentEngine *ConcurrentEngine) Run(seeds ...Request)  {
	output := make(chan ParseResult)
	concurrentEngine.Scheduler.Run()

	for i := 0; i < concurrentEngine.WorkNum; i++ {
		createWork(concurrentEngine.Scheduler.WorkerChan(), output, concurrentEngine.Scheduler)
	}

	for _, seed := range seeds {
		concurrentEngine.Scheduler.Submit(seed)
	}

	for  {
		result := <- output

		for _, item := range result.Items {
			concurrentEngine.ItemChan <- item
		}

		for _, request := range result.Requests {
			_, exists := uniqueUrl[request.Url]

			if exists {
				log.Printf("Got url exist: %s", request.Url)

				continue
			}

			uniqueUrl[request.Url] = true

			concurrentEngine.Scheduler.Submit(request)
		}
	}
}

func createWork(input chan Request, output chan ParseResult, workerReady ReadyNotifier)  {
	go func() {
		for {
			workerReady.WorkerReady(input)
			request := <- input

			log.Printf("Feting: %s", request.Url)
			// 爬取url里面的内容
			body, err := fetcher.Fetch(request.Url)

			if err != nil {
				log.Printf("Fetcher: error " + "fetching url %s:%v", request.Url, err)
			}

			// 根据传递的解析器 解析内容爬取的内容
			parseResult := request.ParserFunc(body)
			// 关闭连接
			body.Close()

			output <- parseResult
		}
	}()
}