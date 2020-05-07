package parse

import (
	"go_crawler/engine"
	"go_crawler/model"
	"regexp"
	"strconv"
)

var (
	autherRe = regexp.MustCompile(`<span class="pl"> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`)
	pressRe  = regexp.MustCompile(`<span class="pl">出版社:</span>([^<]+)<br/>`)
	pagesRe  = regexp.MustCompile(`<span class="pl">页数:</span> ([^<]+)<br/>`)
	priceRe  = regexp.MustCompile(`<span class="pl">定价:</span>([^<]+)<br/>`)
	scoreRe  = regexp.MustCompile(`<strong class="ll rating_num " property="v:average">([^<]+)</strong>`)
	introRe  = regexp.MustCompile(`<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`)
)

func ParseBookDetail(contents []byte, bookname string) engine.ParseResult {
	bookdetail := model.Bookdetails{}

	bookdetail.Bookname = bookname
	bookdetail.Author = ExtraString(contents, autherRe)
	bookdetail.Press = ExtraString(contents, pressRe)
	pages, err := strconv.Atoi(ExtraString(contents, pagesRe))
	if err == nil {
		bookdetail.Pages = pages
	}
	bookdetail.Price = ExtraString(contents, priceRe)
	bookdetail.Score = ExtraString(contents, scoreRe)
	bookdetail.Intro = ExtraString(contents, introRe)

	result := engine.ParseResult{
		Items: []interface{}{bookdetail},
	}

	return result
}

func ExtraString(contents []byte, re *regexp.Regexp) string {
	matches := re.FindSubmatch(contents)
	if len(matches) >= 2 {
		return string(matches[1])
	} else {
		return ""
	}
}
