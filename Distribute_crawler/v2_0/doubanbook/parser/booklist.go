package parser

import (
	"crawler_v2.0/engine"
	"regexp"
)

const BooklistRe = `<a href="([^"]+)" title="([^"]+)"`

func ParseBooklist(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(BooklistRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		url := string(m[1])
		bookname := string(m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseBookDetail(c, url, bookname)
			},
		})
	}

	return result
}
