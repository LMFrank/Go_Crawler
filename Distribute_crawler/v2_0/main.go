package main

import (
	"crawler_v2.0/engine"
	"crawler_v2.0/schedular"
	"crawler_v2.0/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Schedular:   &schedular.SimpleSchedular{},
		WorkerCount: 10,
	}

	//e.Run(engine.Request{
	//	Url:       "http://www.zhenai.com/zhenghun",
	//	ParseFunc: parser.ParseCityList,
	//})
	e.Run(engine.Request{
		Url:       "https://book.douban.com",
		ParseFunc: parser.ParseTag,
	})
}
