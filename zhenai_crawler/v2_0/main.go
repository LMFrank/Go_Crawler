package main

import (
	"cralwer_v2.0/engine"
	"cralwer_v2.0/schedular"
	"cralwer_v2.0/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Schedular:   &schedular.SimpleSchedular{},
		WorkerCount: 10,
	}

	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
