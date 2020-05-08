package main

import (
	"go_crawler/db"
	"go_crawler/engine"
	"go_crawler/parse/zhengai"
	"go_crawler/scheduler"
)

func main() {
	itemsave, err := db.SaveItem()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkCount: 100,
		ItemChan:  itemsave,
	}
	e.Run(engine.Request{
		Url:   "http://www.zhenai.com/zhenghun",
		Parse: zhengai.ParseCity,
	})

}
