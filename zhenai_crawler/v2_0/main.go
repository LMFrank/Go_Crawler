package main

import (
	"cralwer_v2.0/engine"
	"cralwer_v2.0/zhenai/parser"
)

func main() {
	engine.SimpleEngine{}.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
