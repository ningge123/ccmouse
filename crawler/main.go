package main

import (
	"ccmouse/crawler/engine"
	"ccmouse/crawler/persist"
	"ccmouse/crawler/scheduler"
	"ccmouse/crawler/zhenai/parser"
)

func main() {
	index := "dating_profile"
	saver, err := persist.ItemSaver(index)

	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkNum: 10,
		ItemChan: saver,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}