package main

import (
	"crawler_v2.0/distribute/config"
	itemsaverClient "crawler_v2.0/distribute/persist/client"
	workerClient "crawler_v2.0/distribute/worker/client"
	"crawler_v2.0/doubanbook/parser"
	"crawler_v2.0/engine"
	"crawler_v2.0/schedular"
	"fmt"
)

func main() {
	itemChan, err := itemsaverClient.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	processor, err := workerClient.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Schedular:        &schedular.QueuedSchedular{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:    "https://book.douban.com",
		Parser: engine.NewFuncParser(parser.ParseBookTag, config.ParseBookTag),
	})
}
