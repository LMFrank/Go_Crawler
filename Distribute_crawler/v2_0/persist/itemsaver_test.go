package persist

import (
	"context"
	"crawler_v2.0/engine"
	"crawler_v2.0/model"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "https://book.douban.com/subject/1060068",
		Type: "book",
		Id:   "1060068",
		Payload: model.Profile{
			Bookname: "撒哈拉的故事",
			Author:   "三毛",
			Press:    "哈尔滨出版社",
			Pages:    217,
			Price:    "15.80元",
			Score:    9.2,
			Intro:    "三毛作品中最脍炙人口的《撒哈拉的故事》，由12篇精彩动人的散文结合而成，其中《沙漠中的饭店》，是三毛适应荒凉单调的沙漠生活后，重新拾笔的第一篇文字，自此之后，三毛便写出一系列以沙漠为背景的故事，倾倒了全世界的中文读者。",
		},
	}

	// Save expected item
	err := save(expected)
	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.148.130:9200/"),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	// Fetch saved item
	resp, err := client.Get().
		Index("douban").
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", *resp.Source)

	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	// Verify result
	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
