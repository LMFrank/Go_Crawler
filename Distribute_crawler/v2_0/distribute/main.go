package main

import (
	"crawler_v2.0/distribute/persist/client"
	"crawler_v2.0/doubanbook/parser"
	"crawler_v2.0/engine"
	"crawler_v2.0/schedular"
)

func main() {
	itemChan, err := client.ItemSaver(":1234")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Schedular:   &schedular.QueuedSchedular{},
		WorkerCount: 10,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		Url:       "https://book.douban.com",
		ParseFunc: parser.ParseTag,
	})
}
