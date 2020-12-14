package main

import (
	"crawler_v2.0/distribute/config"
	"crawler_v2.0/distribute/rpcsupport"
	"crawler_v2.0/distribute/worker"
	"fmt"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	req := worker.Request{
		Url: "https://book.douban.com/subject/1060068",
		Parser: worker.SerializedParser{
			Name: config.ParseBookDetail,
			Args: "撒哈拉的故事",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
