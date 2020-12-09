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
			ISBN:     9787806398791,
		},
	}

	const index = "douban_test"

	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.148.130:9200/"),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	// Save expected item
	err = Save(client, index, expected)
	if err != nil {
		panic(err)
	}

	// Fetch saved item
	resp, err := client.Get().
		Index(index).
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
