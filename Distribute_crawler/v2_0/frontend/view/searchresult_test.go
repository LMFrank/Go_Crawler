package view

import (
	"os"
	"testing"

	"crawler_v2.0/engine"
	"crawler_v2.0/frontend/model"
	common "crawler_v2.0/model"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView(
		"template.html")

	out, err := os.Create("template.test.html")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url:  "https://book.douban.com/subject/1060068",
		Type: "book",
		Id:   "1060068",
		Payload: common.Profile{
			Bookname: "撒哈拉的故事",
			Author:   "三毛",
			Press:    "哈尔滨出版社",
			Pages:    217,
			Price:    "15.80元",
			Score:    9.2,
			ISBN:     9787806398791,
		},
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)
	if err != nil {
		t.Error(err)
	}

}
