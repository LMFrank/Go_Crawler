package main

import (
	"go_crawler/engine"
	"go_crawler/parse"
)

func main() {

	e := engine.ConcurrentEngine{
		Scheduler: &engine.SimpleSchedular{},
		WorkCount: 100,
	}
	e.Run(engine.Request{
		Url:       "https://book.douban.com",
		ParseFunc: parse.ParseTag,
	})

	//engine.Run(engine.Request{
	//	Url: "https://book.douban.com/tag/%E7%A5%9E%E7%BB%8F%E7%BD%91%E7%BB%9C",
	//	ParseFunc: parse.ParseBooklist,
	//})

	//engine.Run(engine.Request{
	//	Url: "https://book.douban.com/subject/30293801/",
	//	ParseFunc: parse.ParseBookDetail,
	//})
}
