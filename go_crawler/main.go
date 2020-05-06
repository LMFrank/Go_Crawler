package main

import (
	"go_crawler/engine"
	"go_crawler/parse"
)

func main() {
	engine.Run(engine.Request{
		Url:       "https://book.douban.com",
		ParseFunc: parse.ParseContent,
	})
}
