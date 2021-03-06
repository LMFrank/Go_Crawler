package main

import (
	"crawler_v2.0/distribute/config"
	"crawler_v2.0/doubanbook/parser"
	"crawler_v2.0/engine"
	"crawler_v2.0/persist"
	"crawler_v2.0/schedular"
)

func main() {
	itemChan, err := persist.ItemSaver("douban")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Schedular:        &schedular.QueuedSchedular{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url:    "https://book.douban.com",
		Parser: engine.NewFuncParser(parser.ParseBookTag, config.ParseBookTag),
	})
}
