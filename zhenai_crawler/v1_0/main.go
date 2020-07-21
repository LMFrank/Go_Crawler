package main

import (
	"cralwer_v1.0/engine"
	"cralwer_v1.0/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
