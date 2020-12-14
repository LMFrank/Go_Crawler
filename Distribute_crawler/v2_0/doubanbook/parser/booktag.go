package parser

import (
	"crawler_v2.0/distribute/config"
	"crawler_v2.0/engine"
	"regexp"
)

const regexpStr = `<a href="([^"]+)" class="tag">([^"]+)</a>`

func ParseBookTag(content []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(regexpStr)
	matches := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    "https://book.douban.com" + string(m[1]),
			Parser: engine.NewFuncParser(ParseBookList, config.ParseBookList),
		})
	}

	return result
}
