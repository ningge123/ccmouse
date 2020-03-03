package engine

import (
	"log"
	"ccmouse/crawler/fetcher"
)

func Run(requests ...Request)  {
	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		parseResult, err := Work(request)

		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

func Work(request Request) (ParseResult, error)  {
	log.Printf("Feting: %s", request.Url)
	// 爬取url里面的内容
	body, err := fetcher.Fetch(request.Url)

	if err != nil {
		log.Printf("Fetcher: error " + "fetching url %s:%v", request.Url, err)

		return ParseResult{}, err
	}

	// 根据传递的解析器 解析内容爬取的内容
	parseResult := request.ParserFunc(body)
	// 关闭连接
	body.Close()

	return parseResult, nil
}
