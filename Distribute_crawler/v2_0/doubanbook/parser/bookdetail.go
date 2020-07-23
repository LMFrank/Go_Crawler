package parser

import (
	"crawler_v2.0/engine"
	"crawler_v2.0/model"
	"regexp"
	"strconv"
	"strings"
)

var (
	autherRe = regexp.MustCompile(`<span class="pl">作者:</span>[\d\D]*?<a.*?>([^<]+)</a>`)
	pressRe  = regexp.MustCompile(`<span class="pl">出版社:</span>([^<]+)<br/>`)
	pagesRe  = regexp.MustCompile(`<span class="pl">页数:</span> ([^<]+)<br/>`)
	priceRe  = regexp.MustCompile(`<span class="pl">定价:</span>([^<]+)<br/>`)
	scoreRe  = regexp.MustCompile(`<strong class="ll rating_num " property="v:average">([^<]+)</strong>`)
	introRe  = regexp.MustCompile(`<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`)
	idUrlRe  = regexp.MustCompile(`https://book.douban.com/subject/([\d]+)/`)
)

func ParseBookDetail(contents []byte, url string, bookname string) engine.ParseResult {
	profile := model.Profile{}

	profile.Bookname = bookname
	profile.Author = replaceBlank(extraString(contents, autherRe))
	profile.Press = replaceBlank(extraString(contents, pressRe))
	pages, err := strconv.Atoi(extraString(contents, pagesRe))
	if err == nil {
		profile.Pages = pages
	}
	profile.Price = replaceBlank(extraString(contents, priceRe))
	score, err := strconv.ParseFloat(replaceBlank(extraString(contents, scoreRe)), 64)
	if err == nil {
		profile.Score = score
	}
	profile.Intro = extraString(contents, introRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "book",
				Id:      extraString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}

	return result
}

func extraString(contents []byte, re *regexp.Regexp) string {
	matches := re.FindSubmatch(contents)
	if len(matches) >= 2 {
		return string(matches[1])
	} else {
		return ""
	}
}

func replaceBlank(str string) string {
	str = strings.Replace(str, " ", "", -1)
	return strings.Replace(str, "\n", "", -1)
}
