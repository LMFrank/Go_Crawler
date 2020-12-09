package main

import (
	"crawler_v2.0/distribute/config"
	"crawler_v2.0/distribute/rpcsupport"
	"crawler_v2.0/engine"
	"crawler_v2.0/model"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	go serveRpc(host, "test1")
	time.Sleep(5 * time.Second)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	item := engine.Item{
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

	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result is %s, err is %s", result, err)
	}
}
