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
)

func ParseBookDetail(contents []byte, bookname string) engine.ParseResult {
	profile := model.Profile{}

	profile.Bookname = bookname
	profile.Author = ExtraString(contents, autherRe)
	profile.Press = ExtraString(contents, pressRe)
	pages, err := strconv.Atoi(ExtraString(contents, pagesRe))
	if err == nil {
		profile.Pages = pages
	}
	profile.Price = ExtraString(contents, priceRe)
	score, err := strconv.ParseFloat(ExtraString(contents, scoreRe), 64)
	if err == nil {
		profile.Score = score
	}
	profile.Intro = ExtraString(contents, introRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

func ExtraString(contents []byte, re *regexp.Regexp) string {
	matches := re.FindSubmatch(contents)
	if len(matches) >= 2 {
		str := strings.Replace(string(matches[1]), " ", "", -1)
		return strings.Replace(str, "\n", "", -1)
	} else {
		return ""
	}
}
