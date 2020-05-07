package main

import (
	"go_crawler/db"
	"go_crawler/engine"
	"go_crawler/parse"
	"go_crawler/schedular"
)

func main() {

	e := engine.ConcurrentEngine{
		Scheduler: &schedular.QueueSchedular{},
		//Scheduler: &schedular.SimpleSchedular{},
		WorkCount: 100,
		ItemChan:  db.SaveItem(),
	}
	e.Run(engine.Request{
		Url:       "https://book.douban.com",
		ParseFunc: parse.ParseTag,
	})

}
