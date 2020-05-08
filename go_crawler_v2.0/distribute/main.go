package main

import (
	"go_crawler/distribute/client"
	client2 "go_crawler/distribute/work/client"
	"go_crawler/engine"
	"go_crawler/parse/zhengai"
	"go_crawler/scheduler"
)

func main() {

	itemsave, err := client.ItemSave(":1234")

	process, err := client2.CreateProcess()
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueueScheduler{},
		WorkCount:        100,
		ItemChan:         itemsave,
		RequestProcessor: process,
	}

	e.Run(engine.Request{
		Url:   "http://www.zhenai.com/zhenghun",
		Parse: engine.NewFuncparse(zhengai.ParseCityList, "ParseCityList"),
	})
}
