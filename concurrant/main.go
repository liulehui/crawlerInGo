package main

import (
	"github.com/liulehui/crawler/concurrant/config"
	"github.com/liulehui/crawler/concurrant/engine"
	"github.com/liulehui/crawler/concurrant/persist"
	"github.com/liulehui/crawler/concurrant/scheduler"
	"github.com/liulehui/crawler/concurrant/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver(
		config.ElasticIndex)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url: "http://www.starter.url.here",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList),
	})
}
